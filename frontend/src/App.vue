<script lang="ts" setup>
import {useDrvice} from "./store/useDrvice";
import {ref, onMounted} from "vue";
import {ElMessageBox,ElMessage} from "element-plus";

const drviceStore = useDrvice();
const isLoading = ref(true);

onMounted(async () => {
  try {
    await Promise.all([
      drviceStore.getDrviceInfo(),
      drviceStore.getCode()
    ]).then(() => {
      ElMessage({
        message: '设备信息加载完成',
        type: 'success',
        duration: 2000
      })
      isLoading.value = false
    }).catch(() => {
      ElMessageBox.alert(
        '获取设备信息失败，请确保您的电脑有wmic或者powershel3.0及以上版本。',
        '加载失败',
        {
          confirmButtonText: '确定',
          type: 'error'
        });
    })
  } catch (error) {
    console.error('加载数据失败:', error);
  }
});
</script>

<template>
  <div v-if="isLoading" class="loading-container">
    <div class="loading-spinner"></div>
    <p>加载中...</p>
  </div>
  <router-view v-else/>
</template>

<style scoped>
.loading-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100vh;
  width: 100vw;
}

.loading-spinner {
  width: 50px;
  height: 50px;
  border: 4px solid #f3f3f3;
  border-top: 4px solid #409EFF;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}

.loading-container p {
  margin-top: 20px;
  font-size: 16px;
  color: #606266;
}
</style>
