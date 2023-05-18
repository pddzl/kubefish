<template>
  <div class="format">
    <vue-json-pretty :data="data" :deep="1" :showLine="true" />
  </div>
</template>

<script lang="ts" setup>
import { toRefs } from "vue"
import VueJsonPretty from "vue-json-pretty"
import "vue-json-pretty/lib/styles.css"

defineOptions({
  name: "JsonPretty"
})

const props = defineProps({
  data: {
    type: Object,
    default: () => {}
  },
  color: {
    type: String,
    default: () => "dark"
  }
})
const { color } = toRefs(props)
</script>

<style lang="scss" scoped>
.format {
  color: v-bind(color);
}
</style>
