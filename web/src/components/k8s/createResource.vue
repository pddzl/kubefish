<template>
  <div @click="handleClick">
    <el-tooltip effect="dark" :content="tips" placement="bottom">
      <SvgIcon name="plus" />
      <!-- <el-icon><Plus /></el-icon> -->
    </el-tooltip>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue"
import { useRouter } from "vue-router"

const router = useRouter()

const tips = ref("创建资源")

const handleClick = () => {
  router.push({ name: "CreateResource" })
}
</script>

<style lang="scss" scoped>
.svg-icon {
  font-size: 20px;
  &:focus {
    outline: none;
  }
}
</style>
