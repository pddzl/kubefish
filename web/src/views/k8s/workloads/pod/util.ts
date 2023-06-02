// 跳转webssh页面
export const sshPod = (pod: string, namespace: string) => {
  const url = `http://${window.location.host}/#/k8s/pod/webssh?pod=${pod}&namespace=${namespace}`
  const width = 1300
  const height = 800
  // 获取屏幕的宽度和高度
  const screenWidth = window.screen.width
  const screenHeight = window.screen.height

  // 计算新窗口的左上角位置
  const left = (screenWidth - width) / 2
  const top = (screenHeight - height) / 2
  // 构建 features 参数，包括宽度、高度和位置信息
  const features = `width=${width},height=${height},left=${left},top=${top}`
  // window.open(url, "_blank", features)
  window.open(url, "_blank", features)
}
