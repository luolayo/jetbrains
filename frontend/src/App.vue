<script lang="ts" setup>
import {useDrvice} from "./store/useDrvice";
import {ref, onMounted} from "vue";
import {ElMessageBox,ElMessage} from "element-plus";
import {CheckPermissions} from "../wailsjs/go/main/App";
import router from "./router";

const drviceStore = useDrvice();
const isLoading = ref(true);

const MacOSCheck = async () => {
  const userAgent = navigator.userAgent.toLowerCase()
 if (userAgent.includes('mac')) {
   const falg = await CheckPermissions()
   if (!falg) {
     await router.push("/getPermission");
   }
 }
}

onMounted(async () => {
  await MacOSCheck()
  try {
    // 先获取设备信息
    await drviceStore.getDrviceInfo()
  } catch (error) {
    console.error('获取设备信息失败:', error);
    await ElMessageBox.alert(
      '获取设备信息失败，请确保您的电脑有wmic或者PowerShell 3.0及以上版本。',
      '加载失败',
      {
        confirmButtonText: '确定',
        type: 'error'
      });
    return;
  }

  try {
    // 再获取激活码
    await drviceStore.getCode()
    ElMessage({
      message: '设备信息加载完成',
      type: 'success',
      duration: 2000
    })
    isLoading.value = false
  } catch (error) {
    console.error('连接服务器失败:', error);
    await ElMessageBox.alert(
      '连接服务器失败，请重启软件后重试。',
      '加载失败',
      {
        confirmButtonText: '确定',
        type: 'error'
      });
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
