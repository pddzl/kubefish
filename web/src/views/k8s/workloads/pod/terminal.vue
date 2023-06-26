<template>
  <div class="terminal-container">
    <el-card shadow="never" class="search-wrapper">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="容器">
          <el-select v-model="searchInfo.container" clearable placeholder="请选择">
            <el-option v-for="item in containers" :key="item" :label="item" :value="item" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="ArrowRight" @click="onSubmit">进 入</el-button>
        </el-form-item>
      </el-form>
      <div style="height: 100%">
        <div ref="terminalRef" />
      </div>
    </el-card>
  </div>
</template>

<script lang="ts" setup>
// 引入xterm css
import "xterm/css/xterm.css"
// import "xterm/dist/addons/fullscreen/fullscreen.css"
import { Terminal } from "xterm"
// 自适应插件
import { FitAddon } from "xterm-addon-fit"
// 全屏插件
// import * as fullscreen from "xterm/lib/addons/fullscreen/fullscreen"
// web链接插件
// import * as webLinks from 'xterm/lib/addons/webLinks/webLinks'
// websocket插件
import { AttachAddon } from "xterm-addon-attach"
import { ref, reactive, onBeforeUnmount } from "vue"
import { useRoute } from "vue-router"
import { getPodDetailApi } from "@/api/k8s/workloads/pod"
import { useUserStoreHook } from "@/store/modules/user"

// 路由
const route = useRoute()
const pod = route.query.pod as string
const namespace = route.query.namespace as string

// 响应式数据
const containers = ref<string[]>([])
const searchInfo = reactive({ namespace: namespace, pod: pod, container: "" })
const terminalRef = ref(null)
// 获取container
const getPodContainer = async () => {
  const podDetail = await getPodDetailApi({ namespace: namespace, pod: pod })
  if (podDetail.code === 0) {
    podDetail.data.spec.containers.forEach((element: { name: string }) => {
      containers.value.push(element.name)
    })
  }
}
getPodContainer()

// xterm
let terminal: Terminal
let socket: WebSocket

const socketOnOpen = () => {
  socket.onopen = () => {
    if (!terminal) {
      initTerminal()
    }
    // terminal.write('\r')
    // const msg = { op: 'stdin', data: 'export TERM=xterm && clear \r' }
    // ws.send(JSON.stringify(msg))
    // ws.send(msg)
  }
}

const socketOnMessage = () => {
  socket.onmessage = () => {
    // const msg = JSON.parse(event.data)
    // if (msg.op === 'stdout') {
    //   terminal.write(msg.data)
    // } else {
    //   console.log('invalid msg op: ' + msg)
    // }
  }
}

const socketOnError = () => {
  socket.onerror = (event) => {
    console.log("[error] Connection error")
    terminal.write("error: " + event.message)
    terminal.dispose()
  }
}

const socketOnClose = () => {
  socket.onclose = (event) => {
    if (event.wasClean) {
      console.log(`[close] Connection closed cleanly, code=${event.code} reason=${event.reason}`)
    } else {
      console.log("[close] Connection died")
      terminal.writeln("")
    }
    terminal.write("Connection Reset By Peer! Try Refresh.")
  }
}

const initSocket = () => {
  // 打开websocket
  console.log("初始化websocket")
  socket = new WebSocket(
    `ws://${window.location.host}/k8s/pod/getPodWebSSH?namespace=${namespace}&pod=${pod}&container=${searchInfo.container}`,
    [useUserStoreHook().token]
  )
  socketOnClose()
  socketOnOpen()
  socketOnMessage()
  socketOnError()
}
initSocket()

// 初始化终端
const initTerminal = () => {
  terminal = new Terminal({
    rows: 40,
    // cols: 300,
    convertEol: true, // 启用时，光标将设置为下一行的开头
    fontSize: 14,
    cursorBlink: true,
    disableStdin: false
  })
  const attachAddon = new AttachAddon(socket)
  const fitAddon = new FitAddon()
  terminal.loadAddon(attachAddon)
  terminal.loadAddon(fitAddon)
  terminal.open(terminalRef.value)
  setTimeout(() => {
    fitAddon.fit()
  }, 5)
  // console.log(terminal.cols, terminal.rows)
  terminal.focus()
  // terminal.onData(function(data) {
  //   const msg = { op: 'stdin', data: data }
  //   ws.send(JSON.stringify(msg))
  // })
  // terminal.toggleFullScreen(true)
  // terminal.on('data', function(data) {
  //   const msg = { op: 'stdin', data: data }
  //   ws.send(JSON.stringify(msg))
  // })
  // terminal.on('resize', function(size) {
  //   console.log('resize: ' + size)
  //   const msg = { op: 'resize', cols: size.cols, rows: rows }
  //   ws.send(JSON.stringify(msg))
  // })
}

const onSubmit = () => {
  if (searchInfo.container !== "") {
    if (Object.prototype.hasOwnProperty.call(socket, "close")) {
      socket.close()
    }
    initSocket()
  }
}

onBeforeUnmount(() => {
  socketOnClose()
})
</script>

<style lang="scss" scoped>
.terminal-container {
  padding: 20px;
  background-color: rgba(35, 45, 65, 0.763);
  height: 100%;
}
.search-wrapper {
  margin-bottom: 20px;
  :deep(.el-card__body) {
    padding-bottom: 2px;
  }
}
</style>
@/api/k8s/cluster/pod @/api/k8s/workloads/pod
