<template>
  <div class="info-box">
    <div class="row">
      <div class="item">
        <p>名称</p>
        <span class="content">{{ metadata.name }}</span>
      </div>
      <div class="item" v-if="metadata.namespace">
        <p>命名空间</p>
        <span class="content">{{ metadata.namespace }}</span>
      </div>
      <div class="item">
        <p>UID</p>
        <span class="content">{{ metadata.uid }}</span>
      </div>
      <div>
        <p>创建时间</p>
        <span class="content">{{ formatDateTime(metadata.creationTimestamp) }}</span>
      </div>
    </div>
    <div class="row">
      <div class="item">
        <p>标签</p>
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
      <p style="font-size: 13px">注释</p>
      <vue-json-pretty :data="annotationsFormat" :color="'lightcoral'" />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, onBeforeUpdate, toRefs } from "vue"
import VueJsonPretty from "@/components/vueJsonPretty/index.vue"
import { formatDateTime } from "@/utils/index"

defineOptions({
  name: "MetaData"
})

const props = defineProps({
  metadata: {
    type: Object,
    default: () => {}
  }
})

const { metadata } = toRefs(props)

const annotationsFormat = ref({})

// 去掉双引号、反斜杠、换行符
onBeforeUpdate(() => {
  if (metadata.value.annotations) {
    try {
      annotationsFormat.value = JSON.parse(
        JSON.stringify(metadata.value.annotations).replace(/\\"/g, '"').replace(/"\{/g, "{").replace(/\\n"/g, "")
      )
    } catch (err) {
      annotationsFormat.value = metadata.value.annotations
    }
  }
})
</script>
