<template>
  <div class="info-box">
    <div class="row">
      <div v-if="metadata.name" class="item">
        <p>名称</p>
        <span class="content">{{ metadata.name }}</span>
      </div>
      <div v-if="metadata.namespace" class="item">
        <p>命名空间</p>
        <span class="content">{{ metadata.namespace }}</span>
      </div>
      <div v-if="metadata.uid" class="item">
        <p>UID</p>
        <span class="content">{{ metadata.uid }}</span>
      </div>
      <div v-if="metadata.creationTimestamp">
        <p>创建时间</p>
        <span class="content">{{ formatDateTime(metadata.creationTimestamp) }}</span>
      </div>
    </div>
    <div v-if="metadata.labels" class="row">
      <div class="item">
        <p>标签:</p>
        <div class="content">
          <span v-for="(label, index) in metadata.labels" :key="index" class="shadow">
            {{ index }}
            <span v-if="label">:</span>
            {{ label }}
          </span>
        </div>
      </div>
    </div>
    <div v-if="JSON.stringify(annotationsFormat) !== '{}'">
      <p style="font-size: 13px">注释:</p>
      <vue-json-pretty :data="annotationsFormat" :color="'lightcoral'" />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { formatDateTime } from "@/utils/index"
import VueJsonPretty from "@/components/vueJsonPretty/index.vue"

defineOptions({
  name: "MetaData"
})

const props = defineProps({
  metadata: {
    type: Object,
    default: () => {}
  }
})

let annotationsFormat = {}

// 去掉双引号、反斜杠、换行符
const format = () => {
  if (props.metadata.annotations) {
    try {
      annotationsFormat = JSON.parse(
        JSON.stringify(props.metadata.annotations).replace(/\\"/g, '"').replace(/"\{/g, "{").replace(/\\n"/g, "")
      )
    } catch (err) {
      annotationsFormat = props.metadata.annotations
    }
  }
}
format()
</script>
