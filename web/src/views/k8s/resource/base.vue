<template>
  <div class="app-container">
    <el-card>
      <div style="background-color: white; padding: 2px; color: gray">
        <span>请输入资源编排内容, 支持JSON|YAML格式</span>
      </div>
      <div>
        <textarea ref="codeEditor" />
      </div>
      <div style="margin-top: 10px">
        <el-button type="primary" @click="createDynamicResourceFunc">创 建</el-button>
        <el-button @click="resetContent">清 空</el-button>
      </div>
    </el-card>
  </div>
</template>

<script lang="ts" setup>
import { createDynamicResourceApi } from "@/api/k8s/resource/base"
import { onBeforeUnmount, onMounted, ref } from "vue"
import { ElMessage } from "element-plus"
// import yaml from 'js-yaml'
// codemirror基础资源引入
import _CodeMirror from "codemirror"
import "codemirror/lib/codemirror.css"
import "codemirror/lib/codemirror.js"
import "codemirror/mode/javascript/javascript.js"
import "codemirror/mode/yaml/yaml.js"

// 代码错误检查
// require('script-loader!jsonlint')
// import jsonlint from 'jsonlint'
// window.jsonlint = jsonlint
// import 'script-loader!jsonlint'
import "codemirror/addon/lint/lint.css"
import "codemirror/addon/lint/lint"
import "codemirror/addon/lint/json-lint"

import "codemirror/addon/display/autorefresh"
import "codemirror/addon/edit/matchbrackets"

// 折叠资源引入:开始
import "codemirror/addon/fold/foldgutter.css"
import "codemirror/addon/fold/foldcode.js"
import "codemirror/addon/fold/brace-fold.js"
import "codemirror/addon/fold/comment-fold.js"
import "codemirror/addon/fold/indent-fold.js"
import "codemirror/addon/fold/foldgutter.js"
// 折叠资源引入:结束

// 搜索资源引入:开始
import "codemirror/addon/scroll/annotatescrollbar.js"
import "codemirror/addon/search/matchesonscrollbar.js"
import "codemirror/addon/search/match-highlighter.js"
import "codemirror/addon/search/jump-to-line.js"

import "codemirror/addon/dialog/dialog.js"
import "codemirror/addon/dialog/dialog.css"
import "codemirror/addon/search/searchcursor.js"
import "codemirror/addon/search/search.js"
// 搜索资源引入:结束

import { EditorFromTextArea } from "codemirror"

// 启用placeholder
// import 'codemirror/addon/display/placeholder.js'

import "codemirror/addon/selection/active-line.js" // 光标行背景高亮，配置里面也需要styleActiveLine设置为true

const CodeMirror = window.CodeMirror || _CodeMirror

const readOnly = ref(false)
const codeEditor = ref(null)
let editor: EditorFromTextArea | null

const createDynamicResourceFunc = async () => {
  const content = editor?.getValue()
  const res = await createDynamicResourceApi({ content: content })
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "创建成功!"
    })
  }
}

const resetContent = () => {
  editor?.setValue("")
}

onMounted(() => {
  const textareaElement = codeEditor.value
  if (textareaElement) {
    editor = CodeMirror.fromTextArea(textareaElement, {
      // value: modelValue.value,
      mode: "yaml",
      // mime: "text/x-yaml",
      indentWithTabs: false, // 在缩进时，是否需要把 n*tab宽度个空格替换成n个tab字符，默认为false
      smartIndent: true, // 自动缩进，设置是否根据上下文自动缩进（和上一行相同的缩进量）。默认为true
      lineNumbers: true, // 是否在编辑器左侧显示行号
      matchBrackets: true, // 括号匹配
      readOnly: readOnly.value,
      autoRefresh: true,
      lint: true, // window.jsonlint not defined, CodeMirror JSON linting cannot run
      // 启用代码折叠相关功能:开始
      foldGutter: true,
      lineWrapping: true,
      gutters: ["CodeMirror-linenumbers", "CodeMirror-foldgutter", "CodeMirror-lint-markers"],
      // 启用代码折叠相关功能:结束
      styleActiveLine: true // 光标行高亮
    })
    editor.setSize("auto", "400px")
  }
  // editor = CodeMirror.fromTextArea(codeEditor.value, {
  //   // value: modelValue.value,
  //   mode: "yaml",
  //   mime: "text/x-yaml",
  //   indentWithTabs: false, // 在缩进时，是否需要把 n*tab宽度个空格替换成n个tab字符，默认为false
  //   smartIndent: true, // 自动缩进，设置是否根据上下文自动缩进（和上一行相同的缩进量）。默认为true
  //   lineNumbers: true, // 是否在编辑器左侧显示行号
  //   matchBrackets: true, // 括号匹配
  //   readOnly: readOnly.value,
  //   autoRefresh: true,
  //   lint: true, // window.jsonlint not defined, CodeMirror JSON linting cannot run
  //   // 启用代码折叠相关功能:开始
  //   foldGutter: true,
  //   lineWrapping: true,
  //   gutters: ["CodeMirror-linenumbers", "CodeMirror-foldgutter", "CodeMirror-lint-markers"],
  //   // 启用代码折叠相关功能:结束
  //   styleActiveLine: true // 光标行高亮
  // })
})
onBeforeUnmount(() => {
  if (editor !== null) {
    editor.toTextArea()
    editor = null
  }
})
</script>
