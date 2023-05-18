<template>
  <div>
    <div style="position: absolute; top: 15px; right: 20px">
      <el-select v-model="mimeType" size="small">
        <el-option v-for="mine in mimeList" :key="mine" :label="mine" :value="mine" />
      </el-select>
    </div>
    <div>
      <textarea ref="codeEditor" />
    </div>
  </div>
</template>

<script lang="ts">
import { onBeforeUnmount, onMounted, ref, toRefs, watch } from "vue"
import yaml from "js-yaml"
// codemirror基础资源引入
import _CodeMirror from "codemirror"
import "codemirror/lib/codemirror.css"
import "codemirror/lib/codemirror.js"
import "codemirror/mode/javascript/javascript.js"
import "codemirror/mode/yaml/yaml.js"

// 代码错误检查
// import 'script-loader!jsonlint'
// import 'codemirror/addon/lint/lint.css'
// import 'codemirror/addon/lint/lint'
// import 'codemirror/addon/lint/json-lint'

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

// 启用placeholder
// import 'codemirror/addon/display/placeholder.js'

import "codemirror/addon/selection/active-line.js" // 光标行背景高亮，配置里面也需要styleActiveLine设置为true

const CodeMirror = window.CodeMirror || _CodeMirror

export default {
  name: "VueCodeMirror",
  props: {
    modelValue: {
      type: String,
      default: () => ""
    },
    defaultValue: {
      type: String,
      default: ""
    },
    readOnly: {
      type: Boolean,
      default: false
    }
  },
  setup(props, context) {
    const mimeList = ref(["YAML", "JSON"])
    const mimeType = ref("YAML")
    const { modelValue, defaultValue, readOnly } = toRefs(props)
    const codeEditor = ref()
    let editor: any
    watch(modelValue, () => {
      if (editor != null && modelValue.value && modelValue.value !== editor.getValue()) {
        // 触发v-model的双向绑定
        if (mimeType.value === "JSON") {
          editor.setValue(JSON.stringify(JSON.parse(modelValue.value), null, 2))
        } else if (mimeType.value === "YAML") {
          editor.setValue(yaml.dump(JSON.parse(modelValue.value)))
        }
        setTimeout(function () {
          editor.refresh()
        }, 0)
      }
    })
    watch(mimeType, () => {
      if (editor != null && mimeType.value) {
        if (mimeType.value === "JSON") {
          editor.setOption("mode", "javascript")
          editor.setOption("mime", "text/javascript")
          editor.setValue(JSON.stringify(yaml.load(modelValue.value), null, 2))
        } else if (mimeType.value === "YAML") {
          editor.setOption("mode", "yaml")
          editor.setOption("mime", "text/x-yaml")
          editor.setValue(yaml.dump(JSON.parse(modelValue.value)))
        }
        setTimeout(function () {
          editor.refresh()
        }, 0)
      }
    })
    // watch(readOnly, () => {
    //   if (editor != null) {
    //     editor.setOption('readOnly', readOnly.value)
    //   }
    // })
    onMounted(() => {
      editor = CodeMirror.fromTextArea(codeEditor.value, {
        value: modelValue.value,
        // mode: 'yaml',
        // mime: 'text/x-yaml',
        indentWithTabs: false, // 在缩进时，是否需要把 n*tab宽度个空格替换成n个tab字符，默认为false
        smartIndent: true, // 自动缩进，设置是否根据上下文自动缩进（和上一行相同的缩进量）。默认为true
        lineNumbers: true, // 是否在编辑器左侧显示行号
        matchBrackets: true, // 括号匹配
        readOnly: readOnly.value,
        autoRefresh: true,
        // lint: true, // window.jsonlint not defined, CodeMirror JSON linting cannot run
        // 启用代码折叠相关功能:开始
        foldGutter: true,
        lineWrapping: true,
        gutters: ["CodeMirror-linenumbers", "CodeMirror-foldgutter", "CodeMirror-lint-markers"],
        // 启用代码折叠相关功能:结束
        styleActiveLine: true // 光标行高亮
      })
      // 监听编辑器的change事件
      editor.on("change", () => {
        // 触发v-model的双向绑定
        context.emit("update:modelValue", editor.getValue())
      })
      if (defaultValue.value) {
        editor.setValue(defaultValue.value)
      } else {
        if (mimeType.value === "JSON") {
          editor.setOption("mode", "javascript")
          editor.setOption("mime", "text/javascript")
          editor.setValue(JSON.stringify(JSON.parse(modelValue.value), null, 2))
        } else if (mimeType.value === "YAML") {
          editor.setOption("mode", "yaml")
          editor.setOption("mime", "text/x-yaml")
          editor.setValue(yaml.dump(JSON.parse(modelValue.value)))
        }
      }
      editor.setSize("auto", "400px")
    })
    onBeforeUnmount(() => {
      if (editor !== null) {
        editor.toTextArea()
        editor = null
      }
    })
    return {
      codeEditor,
      mimeList,
      mimeType
    }
  }
}
</script>
