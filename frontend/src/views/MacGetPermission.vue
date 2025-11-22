<template>
  <div class="min-h-screen bg-gradient-to-br from-slate-50 to-slate-100 flex items-center justify-center p-4">
    <div class="max-w-2xl w-full">
      <!-- 主卡片 -->
      <div class="bg-white rounded-2xl shadow-xl p-8">
        <!-- 顶部错误提示 -->
        <div class="flex items-center justify-center mb-6">
          <div class="w-20 h-20 bg-red-100 rounded-full flex items-center justify-center">
            <svg class="w-12 h-12 text-red-600" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M13.477 14.89A6 6 0 015.11 6.524l8.367 8.368zm1.414-1.414L6.524 5.11a6 6 0 018.367 8.367zM18 10a8 8 0 11-16 0 8 8 0 0116 0z" clip-rule="evenodd" />
            </svg>
          </div>
        </div>

        <!-- 标题 -->
        <div class="text-center mb-6">
          <h1 class="text-3xl font-bold text-slate-900 mb-2">访问受保护目录失败</h1>
          <p class="text-slate-600">可能缺少完全磁盘访问权限</p>
        </div>

        <!-- 提示信息 -->
        <div class="bg-amber-50 border-2 border-amber-200 rounded-xl p-4 mb-6">
          <div class="flex gap-3">
            <svg class="w-5 h-5 text-amber-600 flex-shrink-0 mt-0.5" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z" clip-rule="evenodd" />
            </svg>
            <div class="flex-1">
              <p class="text-sm font-medium text-amber-900 mb-1">重要提示</p>
              <p class="text-sm text-amber-800">
                为了正常激活 JetBrains 产品，本应用需要访问相关配置文件。请按照以下步骤授予权限。
              </p>
            </div>
          </div>
        </div>

        <!-- 步骤说明 -->
        <div class="mb-6">
          <h2 class="text-lg font-semibold text-slate-900 mb-4">请按以下步骤手动授予权限：</h2>

          <div class="space-y-4">
            <!-- 步骤 1 -->
            <div class="flex gap-4">
              <div class="flex-shrink-0">
                <div class="w-8 h-8 bg-blue-600 text-white rounded-full flex items-center justify-center font-bold">
                  1
                </div>
              </div>
              <div class="flex-1 pt-1">
                <p class="text-slate-700">
                  点击下方按钮（或手动打开<span class="font-semibold">系统设置</span>）
                </p>
              </div>
            </div>

            <!-- 步骤 2 -->
            <div class="flex gap-4">
              <div class="flex-shrink-0">
                <div class="w-8 h-8 bg-blue-600 text-white rounded-full flex items-center justify-center font-bold">
                  2
                </div>
              </div>
              <div class="flex-1 pt-1">
                <p class="text-slate-700">
                  左侧选择<span class="font-semibold">「隐私与安全性」</span> → <span class="font-semibold">「完全磁盘访问权限」</span>
                </p>
              </div>
            </div>

            <!-- 步骤 3 -->
            <div class="flex gap-4">
              <div class="flex-shrink-0">
                <div class="w-8 h-8 bg-blue-600 text-white rounded-full flex items-center justify-center font-bold">
                  3
                </div>
              </div>
              <div class="flex-1 pt-1">
                <p class="text-slate-700">
                  点击右下角的锁 🔓 解锁 → 点击 <span class="font-semibold">+</span> 号添加 Jetbrains到列表中
                </p>
              </div>
            </div>

            <!-- 步骤 4 -->
            <div class="flex gap-4">
              <div class="flex-shrink-0">
                <div class="w-8 h-8 bg-blue-600 text-white rounded-full flex items-center justify-center font-bold">
                  4
                </div>
              </div>
              <div class="flex-1 pt-1">
                <p class="text-slate-700">
                  <span class="font-semibold text-red-600">重启本应用</span>以使权限生效
                </p>
              </div>
            </div>
          </div>
        </div>

        <!-- 操作按钮 -->
        <div class="flex flex-col gap-3">
          <button
            @click="openSystemSettings"
            :disabled="isOpening"
            :class="[
              'px-6 py-4 rounded-xl font-semibold text-lg transition-all shadow-lg',
              isOpening
                ? 'bg-slate-400 cursor-not-allowed'
                : 'bg-gradient-to-r from-blue-600 to-indigo-600 hover:from-blue-700 hover:to-indigo-700 text-white hover:shadow-xl transform hover:scale-105'
            ]"
          >
            <span v-if="!isOpening" class="flex items-center justify-center gap-2">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
              </svg>
              打开系统设置
            </span>
            <span v-else class="flex items-center justify-center gap-3">
              <svg class="w-6 h-6 animate-spin" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              正在打开...
            </span>
          </button>
        </div>

        <!-- 底部提示 -->
        <div class="mt-6 pt-6 border-t border-slate-200">
          <div class="flex items-start gap-3">
            <svg class="w-5 h-5 text-blue-600 flex-shrink-0 mt-0.5" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd" />
            </svg>
            <div class="flex-1">
              <p class="text-sm text-slate-600">
                <span class="font-semibold">提示：</span>授予权限后，请务必重启本应用才能生效。如果问题仍然存在，请检查是否已正确添加应用到「完全磁盘访问权限」列表中。
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { BrowserOpenURL } from '../../wailsjs/runtime'

const isOpening = ref(false)

// 打开系统设置
const openSystemSettings = async () => {
  isOpening.value = true

  try {
    BrowserOpenURL('x-apple.systempreferences:com.apple.preference.security?Privacy_AllFiles')
  } catch (error) {
    console.error('打开系统设置失败:', error)
    try {
      BrowserOpenURL('x-apple.systempreferences:')
    } catch (err) {
      console.error('打开系统设置主页失败:', err)
      alert('无法自动打开系统设置，请手动打开：系统设置 → 隐私与安全性 → 完全磁盘访问权限')
    }
  } finally {
    // 延迟重置按钮状态
    setTimeout(() => {
      isOpening.value = false
    }, 1000)
  }
}
</script>

<style scoped>
/* 可以添加额外的自定义样式 */
</style>