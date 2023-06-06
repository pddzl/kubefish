package k8s

import (
	"bytes"
	"context"
	"github.com/gorilla/websocket"
	"github.com/pddzl/kubefish/server/global"
	modelK8s "github.com/pddzl/kubefish/server/model/k8s"
	"io"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
)

type PodService struct{}

// GetPods 获取集群pods
func (ps *PodService) GetPods(namespace string, label string, page int, pageSize int) ([]modelK8s.PodBrief, int, error) {
	// 获取pod list
	list, err := global.KF_K8S_Client.CoreV1().Pods(namespace).List(context.TODO(), metaV1.ListOptions{LabelSelector: label})
	if err != nil {
		return nil, 0, err
	}

	var podList coreV1.PodList
	// 分页
	end := pageSize * page
	offset := pageSize * (page - 1)
	total := len(list.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		podList.Items = list.Items[offset:]
	} else {
		podList.Items = list.Items[offset:end]
	}

	// 处理list数据
	var podBriefList []modelK8s.PodBrief
	for _, pod := range podList.Items {
		var podBrief modelK8s.PodBrief
		podBrief.Name = pod.Name
		podBrief.Namespace = pod.Namespace
		podBrief.Node = pod.Spec.NodeName
		podBrief.Status = string(pod.Status.Phase)
		podBrief.CreationTimestamp = pod.CreationTimestamp.Time
		// append
		podBriefList = append(podBriefList, podBrief)
	}
	return podBriefList, total, nil
}

// GetPodDetail 获取指定pod详情
func (ps *PodService) GetPodDetail(namespace string, pod string) (*modelK8s.PodDetail, error) {
	podDetailRaw, err := global.KF_K8S_Client.CoreV1().Pods(namespace).Get(context.TODO(), pod, metaV1.GetOptions{})
	if err != nil {
		return nil, nil
	}

	var podDetail modelK8s.PodDetail

	// metadata
	podDetail.MetaData = modelK8s.NewObjectMeta(podDetailRaw.ObjectMeta)

	// spec
	podDetail.Spec.NodeName = podDetailRaw.Spec.NodeName
	podDetail.Spec.RestartPolicy = string(podDetailRaw.Spec.RestartPolicy)
	podDetail.Spec.ServiceAccountName = podDetailRaw.Spec.ServiceAccountName

	// spec -> initContainer | container
	podDetail.Spec.InitContainers = extractContainerInfo(podDetailRaw.Spec.InitContainers, podDetailRaw)
	podDetail.Spec.Containers = extractContainerInfo(podDetailRaw.Spec.Containers, podDetailRaw)

	// status
	podDetail.Status.Phase = string(podDetailRaw.Status.Phase)
	podDetail.Status.PodIP = podDetailRaw.Status.PodIP
	podDetail.Status.QOSClass = string(podDetailRaw.Status.QOSClass)
	podDetail.Status.Conditions = podDetailRaw.Status.Conditions

	return &podDetail, err
}

func extractContainerInfo(containerList []coreV1.Container, pod *coreV1.Pod) (containers []modelK8s.Container) {
	for _, value := range containerList {
		var container modelK8s.Container
		container.Name = value.Name
		container.Image = value.Image
		container.Command = value.Command
		container.Args = value.Args
		container.SecurityContext = value.SecurityContext
		container.LivenessProbe = value.LivenessProbe
		container.ReadinessProbe = value.ReadinessProbe
		container.StartupProbe = value.StartupProbe
		// volume
		for _, volume := range value.VolumeMounts {
			var volumeMount modelK8s.VolumeMount
			volumeMount.Name = volume.Name
			volumeMount.ReadOnly = volume.ReadOnly
			volumeMount.MountPath = volume.MountPath
			for _, specV := range pod.Spec.Volumes {
				if volume.Name == specV.Name {
					if specV.VolumeSource.ConfigMap != nil {
						volumeMount.VolumeType = "configMap"
					} else if specV.VolumeSource.Secret != nil {
						volumeMount.VolumeType = "secret"
					} else if specV.VolumeSource.EmptyDir != nil {
						volumeMount.VolumeType = "emptyDir"
					} else if specV.VolumeSource.Projected != nil {
						volumeMount.VolumeType = "projected"
					} else if specV.VolumeSource.HostPath != nil {
						volumeMount.VolumeType = "hostPath"
					} // ...待续
				}
			}
			container.VolumeMounts = append(container.VolumeMounts, volumeMount)
		}
		// status
		for _, status := range pod.Status.ContainerStatuses {
			if status.Name == value.Name {
				container.Status = &status
			}
		}
		containers = append(containers, container)
	}
	return
}

// GetPodLog 获取pod日志
func (ps *PodService) GetPodLog(namespace string, pod string, container string, lines int64, follow bool) (string, error) {
	req := global.KF_K8S_Client.CoreV1().Pods(namespace).GetLogs(pod, &coreV1.PodLogOptions{Container: container, Follow: follow, TailLines: &lines})

	readCloser, err := req.Stream(context.Background())
	if err != nil {
		return "", err
	}
	defer readCloser.Close()

	// 读取日志内容
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, readCloser)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// DeletePod 删除pod
func (ps *PodService) DeletePod(namespace string, pod string) error {
	return global.KF_K8S_Client.CoreV1().Pods(namespace).Delete(context.TODO(), pod, metaV1.DeleteOptions{})
}

// --- pod WebSSH ---

type TerminalSession struct {
	wsConn   *websocket.Conn
	sizeChan chan remotecommand.TerminalSize
	doneChan chan struct{}
}

func (ps *PodService) NewTerminalSession(conn *websocket.Conn) (*TerminalSession, error) {
	session := &TerminalSession{
		wsConn:   conn,
		sizeChan: make(chan remotecommand.TerminalSize),
		doneChan: make(chan struct{}),
	}
	return session, nil
}

const END_OF_TRANSMISSION = "\u0004"

// PtyHandler is what remoteCommand expects from a pty
type PtyHandler interface {
	io.Reader
	io.Writer
	remotecommand.TerminalSizeQueue
}

// GetPodWebSSH pod webSSH
func (ps *PodService) GetPodWebSSH(namespace string, podName string, containerName string, cfg *rest.Config, ptyHandler PtyHandler) error {
	req := global.KF_K8S_Client.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(podName).
		Namespace(namespace).
		SubResource("exec")

	req.VersionedParams(&coreV1.PodExecOptions{
		Container: containerName,
		Command:   []string{"sh", "-c", "test -f /bin/bash && bash || sh"},
		Stdin:     true,
		Stdout:    true,
		Stderr:    true,
		TTY:       true,
	}, scheme.ParameterCodec)

	exec, err := remotecommand.NewSPDYExecutor(cfg, "POST", req.URL())
	if err != nil {
		return err
	}

	err = exec.StreamWithContext(context.Background(), remotecommand.StreamOptions{
		Stdin:             ptyHandler,
		Stdout:            ptyHandler,
		Stderr:            ptyHandler,
		TerminalSizeQueue: ptyHandler,
		Tty:               true,
	})
	if err != nil {
		return err
	}

	return nil
}

// TerminalMessage is the messaging protocol between ShellController and TerminalSession.
//
// OP      DIRECTION  FIELD(S) USED  DESCRIPTION
// ---------------------------------------------------------------------
// bind    fe->be     SessionID      Id sent back from TerminalResponse
// stdin   fe->be     Data           Keystrokes/paste buffer
// resize  fe->be     Rows, Cols     New terminal size
// stdout  be->fe     Data           Output from the process
// toast   be->fe     Data           OOB message to be shown to the user
type TerminalMessage struct {
	Op   string `json:"op"`
	Data string `json:"data"`
	//SessionID string `json:"sessionID"`
	Rows uint16 `json:"rows"`
	Cols uint16 `json:"cols"`
}

// Next TerminalSize handles pty->process resize events
// Called in a loop from remotecommand as long as the process is running
func (t *TerminalSession) Next() *remotecommand.TerminalSize {
	select {
	case size := <-t.sizeChan:
		return &size
	case <-t.doneChan:
		return nil
	}
}

// Read handles pty->process messages (stdin, resize)
// Called in a loop from remotecommand as long as the process is running
func (t *TerminalSession) Read(p []byte) (int, error) {
	_, message, err := t.wsConn.ReadMessage()
	//global.KF_LOG.Info("message", zap.String("ms", string(message)))
	if err != nil {
		return copy(p, END_OF_TRANSMISSION), err
	}
	return copy(p, message), nil
	//var msg TerminalMessage
	//if err := json.Unmarshal([]byte(message), &msg); err != nil {
	//	global.KF_LOG.Error("json", zap.Error(err))
	//	return copy(p, END_OF_TRANSMISSION), err
	//}
	//switch msg.Op {
	//case "stdin":
	//	global.KF_LOG.Info("stdin debug")
	//	return copy(p, msg.Data), nil
	//case "resize":
	//	t.sizeChan <- remotecommand.TerminalSize{Width: msg.Cols, Height: msg.Rows}
	//	return 0, nil
	//default:
	//	global.KF_LOG.Error(fmt.Sprintf("unknown message type '%s'", msg.Op))
	//	return copy(p, END_OF_TRANSMISSION), fmt.Errorf("unknown message type '%s'", msg.Op)
	//}
}

// Write called from remoteCommand whenever there is any output
func (t *TerminalSession) Write(p []byte) (int, error) {
	//msg, err := json.Marshal(TerminalMessage{
	//	Op:   "stdout",
	//	Data: string(p),
	//})
	//if err != nil {
	//	return 0, err
	//}
	//if err := t.wsConn.WriteMessage(websocket.TextMessage, msg); err != nil {
	//	return 0, err
	//}
	if err := t.wsConn.WriteMessage(websocket.TextMessage, p); err != nil {
		return 0, err
	}
	return len(p), nil
}

// Close session
func (t *TerminalSession) Close() error {
	close(t.doneChan)
	return t.wsConn.Close()
}
