<template>
  <div class="min-h-screen bg-gradient-to-br from-slate-50 to-slate-100">
    <!-- 进入页面弹框提示 -->
    <div v-if="showWelcomeDialog" class="fixed inset-0 bg-black/50 flex items-center justify-center z-50">
      <div class="bg-white rounded-2xl shadow-2xl max-w-lg mx-4 overflow-hidden">
        <div class="bg-gradient-to-r from-amber-500 to-orange-500 p-6">
          <div class="flex items-center gap-3">
            <div class="w-12 h-12 bg-white/20 rounded-xl flex items-center justify-center">
              <svg class="w-7 h-7 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
              </svg>
            </div>
            <h2 class="text-xl font-bold text-white">温馨提示</h2>
          </div>
        </div>
        <div class="p-6">
          <p class="text-slate-700 leading-relaxed mb-4">
            本下载功能为<span class="font-semibold text-amber-600">附赠功能</span>，如遇到任何问题（网络问题、下载失败、安装失败等），请直接前往 <span class="font-semibold text-blue-600">JetBrains 官网</span> 自行下载安装即可。
          </p>
          <p class="text-slate-500 text-sm">
            官网地址：https://www.jetbrains.com/
          </p>
        </div>
        <div class="px-6 pb-6">
          <button
            @click="showWelcomeDialog = false"
            class="w-full py-3 bg-gradient-to-r from-amber-500 to-orange-500 text-white font-medium rounded-xl hover:from-amber-600 hover:to-orange-600 transition-all shadow-lg hover:shadow-xl"
          >
            我已知晓，继续使用
          </button>
        </div>
      </div>
    </div>

    <div class="container mx-auto px-4 py-8 max-w-6xl">
      <!-- 返回按钮 -->
      <button
        @click="goBack"
        class="mb-6 flex items-center gap-2 text-slate-600 hover:text-slate-900 transition-colors"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"/>
        </svg>
        返回首页
      </button>

      <!-- 页面标题 -->
      <div class="text-center mb-8">
        <h1 class="text-4xl font-bold text-slate-900 mb-2">下载 JetBrains 软件</h1>
        <p class="text-slate-600">选择您需要的产品、版本和操作系统</p>
      </div>

      <!-- 长提示卡片 -->
      <div class="bg-gradient-to-r from-blue-50 to-indigo-50 border-2 border-blue-200 rounded-2xl p-6 mb-6 mx-18">
        <div class="flex gap-4">
          <div class="flex-shrink-0">
            <div class="w-12 h-12 bg-blue-600 rounded-xl flex items-center justify-center">
              <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </div>
          </div>
          <div class="flex-1">
            <h3 class="text-lg font-semibold text-slate-900 mb-2">下载须知</h3>
            <ul class="space-y-2 text-sm text-slate-700">
              <li class="flex items-start gap-2">
                <svg class="w-5 h-5 text-blue-600 flex-shrink-0 mt-0.5" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                </svg>
                <span>本功能属于附赠功能，所以能不能用我都不会去售后，并且本功能一切来源于官网，所以你官网下载不了本软件也无法下载安装</span>
              </li>
              <li class="flex items-start gap-2">
                <svg class="w-5 h-5 text-blue-600 flex-shrink-0 mt-0.5" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                </svg>
                <span>Windows用户和Linux请注意选择的安装目录就是软件下载安装包和安装软件的地方，软件安装后请打开一次关闭后再使用激活功能</span>
              </li>
              <li class="flex items-start gap-2">
                <svg class="w-5 h-5 text-blue-600 flex-shrink-0 mt-0.5" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                </svg>
                <span>目前本软件过滤了一些乱七八糟的软件，以下的列表里的软件都是常用且能正常激活的</span>
              </li>
              <li class="flex items-start gap-2">
                <svg class="w-5 h-5 text-blue-600 flex-shrink-0 mt-0.5" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                </svg>
                <span>如遇到任何下载或安装问题，请直接前往 JetBrains 官网自行下载安装即可</span>
              </li>
            </ul>
          </div>
        </div>
      </div>

      <!-- 产品选择 -->
      <div class="bg-white rounded-2xl shadow-sm p-6 mb-6">
        <h2 class="text-lg font-semibold text-slate-900 mb-4">选择产品</h2>
        <div class="grid grid-cols-2 md:grid-cols-4 gap-3">
          <button
            v-for="product in products"
            :key="product.id"
            @click="selectedProduct = product.id; loadVersions()"
            :class="[
              'p-4 rounded-xl border-2 transition-all duration-200 flex flex-col items-center gap-3',
              selectedProduct === product.id
                ? 'border-blue-500 bg-blue-50 shadow-md'
                : 'border-slate-200 hover:border-slate-300 hover:shadow-sm'
            ]"
          >
            <img class="w-18 h-18" :src="product.icon" alt="">
            <!--            <div class="text-2xl mb-2">{{ product.icon }}</div>-->
            <div class="font-medium text-slate-900 text-sm">{{ product.name }}</div>
          </button>
        </div>
      </div>

      <!-- 版本和系统选择 -->
      <div class="bg-white rounded-2xl shadow-sm p-6 mb-6">
        <div class="grid md:grid-cols-2 gap-6">
          <!-- 版本选择 -->
          <div>
            <h2 class="text-lg font-semibold text-slate-900 mb-4">选择版本</h2>

            <!-- 版本类型切换 -->
            <div class="flex gap-2 mb-4">
              <button
                @click="versionType = 'latest'"
                :class="[
                  'flex-1 py-2 px-4 rounded-lg font-medium transition-all',
                  versionType === 'latest'
                    ? 'bg-blue-500 text-white shadow-md'
                    : 'bg-slate-100 text-slate-600 hover:bg-slate-200'
                ]"
              >
                最新版本
              </button>
              <button
                @click="versionType = 'history';"
                :class="[
                  'flex-1 py-2 px-4 rounded-lg font-medium transition-all',
                  versionType === 'history'
                    ? 'bg-blue-500 text-white shadow-md'
                    : 'bg-slate-100 text-slate-600 hover:bg-slate-200'
                ]"
              >
                历史版本
              </button>
            </div>


            <div class="relative">
              <div v-if="loadingVersions" class="flex items-center justify-center py-8">
                <div class="flex flex-col items-center gap-3">
                  <div class="w-10 h-10 border-4 border-blue-500 border-t-transparent rounded-full animate-spin"></div>
                  <p class="text-sm text-slate-600">加载版本列表中...</p>
                </div>
              </div>
              <!-- 历史版本选择器 -->
              <div v-else-if="versionType === 'history' && !loadingVersions">
                <button
                  @click="showVersionDropdown = !showVersionDropdown"
                  class="w-full flex items-center justify-between px-4 py-3 bg-slate-50 border-2 border-slate-200 rounded-xl hover:border-slate-300 transition-colors"
                >
                  <span class="font-medium text-slate-900">{{ selectedVersion }}</span>
                  <svg
                    :class="['w-5 h-5 text-slate-600 transition-transform', showVersionDropdown && 'rotate-180']"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/>
                  </svg>
                </button>

                <!-- 下拉列表 -->
                <transition
                  enter-active-class="transition ease-out duration-200"
                  enter-from-class="opacity-0 translate-y-1"
                  enter-to-class="opacity-100 translate-y-0"
                  leave-active-class="transition ease-in duration-150"
                  leave-from-class="opacity-100 translate-y-0"
                  leave-to-class="opacity-0 translate-y-1"
                >
                  <div
                    v-if="showVersionDropdown"
                    class="absolute z-10 w-full mt-2 bg-white border-2 border-slate-200 rounded-xl shadow-lg max-h-64 overflow-y-auto"
                  >
                    <button
                      v-for="version in availableVersions"
                      :key="version"
                      @click="selectVersion(version)"
                      :class="[
                        'w-full px-4 py-3 text-left hover:bg-slate-50 transition-colors flex items-center justify-between',
                        selectedVersion === version && 'bg-blue-50 text-blue-600 font-medium'
                      ]"
                    >
                      <span>{{ version }}</span>
                      <svg
                        v-if="selectedVersion === version"
                        class="w-5 h-5"
                        fill="currentColor"
                        viewBox="0 0 20 20"
                      >
                        <path fill-rule="evenodd"
                              d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                              clip-rule="evenodd"/>
                      </svg>
                    </button>
                  </div>
                </transition>
              </div>

              <!-- 最新版本显示 -->
              <div v-else-if="versionType === 'latest' && !loadingVersions"
                   class="px-4 py-3 bg-gradient-to-r from-blue-50 to-indigo-50 border-2 border-blue-200 rounded-xl">
                <div class="flex items-center gap-2">
                  <svg class="w-5 h-5 text-blue-600" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd"
                          d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                          clip-rule="evenodd"/>
                  </svg>
                  <span class="font-semibold text-blue-900">{{ selectedVersion }}</span>
                  <span class="text-xs bg-blue-600 text-white px-2 py-0.5 rounded-full">最新</span>
                </div>
              </div>
            </div>
          </div>

          <!-- 系统选择 -->
          <div>
            <h2 class="text-lg font-semibold text-slate-900 mb-4">选择操作系统</h2>
            <div class="space-y-2">
              <button
                v-for="os in operatingSystems"
                :key="os.id"
                @click="selectedOS = os.id"
                :class="[
                  'w-full flex items-center gap-3 p-3 rounded-xl border-2 transition-all',
                  selectedOS === os.id
                    ? 'border-blue-500 bg-blue-50 shadow-md'
                    : 'border-slate-200 hover:border-slate-300 hover:shadow-sm'
                ]"
              >
                <span class="text-2xl">{{ os.icon }}</span>
                <span class="font-medium text-slate-900">{{ os.name }}</span>
                <span v-if="os.id === detectedOS"
                      class="ml-auto text-xs bg-green-100 text-green-700 px-2 py-1 rounded-full">
                  已检测
                </span>
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- 保存目录选择区域 -->
      <div class="bg-white rounded-2xl shadow-sm p-6 mb-6">
        <h2 class="text-lg font-semibold text-slate-900 mb-4">安装目录</h2>
        <p class="text-sm text-slate-600 mb-4">选择软件的安装位置，下载完成后将自动安装并删除安装包</p>
        <div class="flex gap-3">
          <button
            @click="selectDir"
            class="px-6 py-2.5 bg-purple-600 text-white rounded-lg hover:bg-purple-700 transition-colors font-medium shadow-sm hover:shadow-md"
          >
            选择目录
          </button>
          <div class="flex-1 px-4 py-2.5 bg-slate-50 border-2 border-slate-200 rounded-lg text-slate-700 truncate flex items-center">
            {{ selectedDir || '未选择目录' }}
          </div>
        </div>
      </div>

      <!-- 下载选项 -->
      <div class="bg-white rounded-2xl shadow-sm p-6 mb-6">
        <h2 class="text-lg font-semibold text-slate-900 mb-4">可用下载</h2>
        <div class="space-y-3">
          <div
            v-for="option in currentDownloadOptions"
            :key="option.link"
            class="flex items-center justify-between p-4 border-2 border-slate-200 rounded-xl hover:border-blue-300 hover:shadow-md transition-all group"
          >
            <div class="flex-1">
              <div class="flex items-center gap-2 mb-1">
                <h3 class="font-semibold text-slate-900">{{ option.name }}</h3>
                <span v-if="option.recommended" class="text-xs bg-blue-600 text-white px-2 py-0.5 rounded-full">
                  推荐
                </span>
              </div>
              <div class="flex items-center gap-4 text-sm text-slate-600">
                <span>{{ option.size }}</span>
                <span>{{ option.format }}</span>
              </div>
            </div>
            <button
              @click="downloadFile(option.link)"
              class="px-6 py-2.5 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors font-medium shadow-sm hover:shadow-md"
            >
              下载
            </button>
          </div>
        </div>
      </div>


      <!-- 下载安装列表区域 -->
      <div class="bg-white rounded-2xl shadow-sm p-6">
        <h2 class="text-lg font-semibold text-slate-900 mb-4">下载安装列表</h2>

        <div v-if="downloadList.length === 0" class="text-center py-8 text-slate-500">
          <svg class="w-16 h-16 mx-auto mb-3 text-slate-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
          </svg>
          <p>暂无待下载文件，点击"加入列表"添加文件</p>
        </div>

        <div v-else class="space-y-4">
          <div
            v-for="item in downloadList"
            :key="item.filePath"
            class="p-4 border-2 border-slate-200 rounded-xl space-y-3"
          >
            <!-- 文件信息和移除按钮 -->
            <div class="flex items-center justify-between">
              <div>
                <h3 class="font-semibold text-slate-900">{{ item.name }}</h3>
                <p class="text-sm text-slate-600">{{ item.size }}</p>
              </div>
            </div>

            <!-- 下载进度条 -->
            <div v-if="item.status !== 'pending'" class="space-y-2">
              <div class="flex items-center justify-between text-sm">
                <span class="text-slate-600">下载进度</span>
                <span class="font-medium text-blue-600">{{ item.downloadProgress }}%</span>
              </div>
              <div class="w-full h-2.5 bg-slate-200 rounded-full overflow-hidden">
                <div
                  class="h-full bg-blue-600 transition-all duration-300"
                  :style="{ width: item.downloadProgress + '%' }"
                ></div>
              </div>
            </div>

            <!-- 安装进度条 -->
            <div v-if="item.status === 'installing' || item.status === 'installed'" class="space-y-2">
              <div class="flex items-center justify-between text-sm">
                <span class="text-slate-600">安装进度</span>
                <span class="font-medium text-green-600">{{ item.installProgress }}%</span>
              </div>
              <div class="w-full h-2.5 bg-slate-200 rounded-full overflow-hidden">
                <div
                  class="h-full bg-green-600 transition-all duration-300"
                  :style="{ width: item.installProgress + '%' }"
                ></div>
              </div>
            </div>

            <!-- 状态标识 -->
            <div class="flex items-center gap-2">
              <span
                :class="[
                  'px-3 py-1 rounded-full text-xs font-medium',
                  item.status === 'pending' && 'bg-slate-100 text-slate-700',
                  item.status === 'downloading' && 'bg-blue-100 text-blue-700',
                  item.status === 'installing' && 'bg-purple-100 text-purple-700',
                  item.status === 'installed' && 'bg-green-100 text-green-700'
                ]"
              >
                {{ getStatusText(item.status) }}
              </span>
              <svg
                v-if="item.status === 'downloading' || item.status === 'installing'"
                class="w-4 h-4 text-blue-600 animate-spin"
                fill="none"
                viewBox="0 0 24 24"
              >
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, onUnmounted, reactive, ref, watch} from 'vue'
import {getVersion} from '../api/download'
import type {DownloadData, DownloadOption, ReleaseData} from '../type'
import {ElMessage} from 'element-plus'
import {CheckDownloadURL, DownloadAndInstall, SelectDirectory} from '../../wailsjs/go/main/App'
import {EventsOff, EventsOn} from '../../wailsjs/runtime/runtime'

// 产品列表
const products = [
  {
    id: 'IIU',
    name: 'IntelliJ IDEA',
    icon: 'https://resources.jetbrains.com.cn/storage/products/intellij-idea/img/meta/intellij-idea_logo_300x300.png'
  },
  {
    id: 'PCP',
    name: 'PyCharm',
    icon: 'https://resources.jetbrains.com.cn/storage/products/pycharm/img/meta/pycharm_logo_300x300.png'
  },
  {
    id: 'WS',
    name: 'WebStorm',
    icon: 'https://resources.jetbrains.com.cn/storage/products/webstorm/img/meta/webstorm_logo_300x300.png'
  },
  {
    id: 'PS',
    name: 'PhpStorm',
    icon: 'https://resources.jetbrains.com.cn/storage/products/phpstorm/img/meta/phpstorm_logo_300x300.png'
  },
  {
    id: 'GO',
    name: 'GoLand',
    icon: 'https://resources.jetbrains.com.cn/storage/products/goland/img/meta/goland_logo_300x300.png'
  },
  {
    id: 'RD',
    name: 'Rider',
    icon: 'https://resources.jetbrains.com.cn/storage/products/rider/img/meta/rider_logo_300x300.png'
  },
  {
    id: 'CL',
    name: 'CLion',
    icon: 'https://resources.jetbrains.com.cn/storage/products/clion/img/meta/clion_logo_300x300.png'
  },
  {
    id: 'DG',
    name: 'DataGrip',
    icon: 'https://resources.jetbrains.com.cn/storage/products/datagrip/img/meta/datagrip_logo_300x300.png'
  },
  {
    id: 'RM',
    name: 'RubyMine',
    icon: 'https://resources.jetbrains.com.cn/storage/products/rubymine/img/meta/rubymine_logo_300x300.png'
  },
  {
    id: 'AC',
    name: 'AppCode',
    icon: 'https://resources.jetbrains.com.cn/storage/products/appcode/img/meta/appcode_logo_300x300.png'
  }
]

// 操作系统列表
const operatingSystems = [
  { id: 'windows', name: 'Windows', icon: '🪟' },
  { id: 'mac', name: 'macOS', icon: '🍎' },
  { id: 'linux', name: 'Linux', icon: '🐧' }
]

// 状态
const selectedProduct = ref('IIU')
const versionType = ref<'latest' | 'history'>('latest')
const selectedVersion = ref('')
const selectedOS = ref('windows')
const detectedOS = ref('windows')
const showVersionDropdown = ref(false)
const loadingVersions = ref(false)
const availableVersions = ref<string[]>([])
// 保存选择的目录路径（安装目录）
const selectedDir = ref('')
// 欢迎弹框状态
const showWelcomeDialog = ref(true)
// 系统默认下载目录
const defaultDownloadDir = ref('')

// 下载安装列表状态
const downloadList = ref<Array<{
  name: string
  link: string
  size: string
  filePath: string
  status: 'pending' | 'downloading' | 'installing' | 'installed'
  downloadProgress: number
  installProgress: number
}>>([])

// 存储所有版本数据
let allReleaseData: ReleaseData[] = []

// 下载选项数据
const downloadData = reactive<{
  windows: DownloadOption[]
  mac: DownloadOption[]
  linux: DownloadOption[]
}>({
  windows: [],
  mac: [],
  linux: []
})

const addToDownloadList = (link: string, name: string, size: string, filePath: string) => {
  // 使用 filePath 作为唯一标识检查是否已存在
  const exists = downloadList.value.some(item => item.filePath === filePath)
  if (exists) {
    ElMessage.warning('该文件已在列表中')
    return false
  }

  downloadList.value.push({
    name,
    link,
    size,
    filePath,
    status: 'pending',
    downloadProgress: 0,
    installProgress: 0
  })
  ElMessage.success('已添加到下载列表')
  return true
}


const getStatusText = (status: string): string => {
  const statusMap: Record<string, string> = {
    'pending': '等待中',
    'downloading': '下载中',
    'installing': '安装中',
    'installed': '已完成'
  }
  return statusMap[status] || '未知'
}


// 格式化文件大小
const formatSize = (bytes: number): string => {
  const gb = bytes / (1024 * 1024 * 1024)
  return `${gb.toFixed(2)} GB`
}

// 获取文件格式
const getFileFormat = (url: string): string => {
  if (url.endsWith('.exe')) return '.exe'
  if (url.endsWith('.dmg')) return '.dmg'
  if (url.endsWith('.tar.gz')) return '.tar.gz'
  if (url.endsWith('.zip')) return '.zip'
  return ''
}

// 设置下载数据
const setDownloadData = () => {
  const release = allReleaseData.find(item => item.version === selectedVersion.value)
  if (!release) {
    ElMessage.error('未找到该版本的数据')
    return
  }

  const downloads = release.downloads

  // Windows 下载选项
  downloadData.windows = []
  if (downloads.windows) {
    downloadData.windows.push({
      name: 'Windows (x64)',
      size: formatSize(downloads.windows.size),
      format: getFileFormat(downloads.windows.link),
      link: downloads.windows.link,
      recommended: true
    })
  }
  if (downloads.windowsARM64) {
    downloadData.windows.push({
      name: 'Windows (ARM64)',
      size: formatSize(downloads.windowsARM64.size),
      format: getFileFormat(downloads.windowsARM64.link),
      link: downloads.windowsARM64.link,
      recommended: false
    })
  }
  if (downloads.windowsZip) {
    downloadData.windows.push({
      name: 'Windows (ZIP)',
      size: formatSize(downloads.windowsZip.size),
      format: getFileFormat(downloads.windowsZip.link),
      link: downloads.windowsZip.link,
      recommended: false
    })
  }

  // macOS 下载选项
  downloadData.mac = []
  if (downloads.macM1) {
    downloadData.mac.push({
      name: 'macOS (Apple Silicon)',
      size: formatSize(downloads.macM1.size),
      format: getFileFormat(downloads.macM1.link),
      link: downloads.macM1.link,
      recommended: true
    })
  }
  if (downloads.mac) {
    downloadData.mac.push({
      name: 'macOS (Intel)',
      size: formatSize(downloads.mac.size),
      format: getFileFormat(downloads.mac.link),
      link: downloads.mac.link,
      recommended: !downloads.macM1
    })
  }

  // Linux 下载选项
  downloadData.linux = []
  if (downloads.linux) {
    downloadData.linux.push({
      name: 'Linux (x64)',
      size: formatSize(downloads.linux.size),
      format: getFileFormat(downloads.linux.link),
      link: downloads.linux.link,
      recommended: true
    })
  }
  if (downloads.linuxARM64) {
    downloadData.linux.push({
      name: 'Linux (ARM64)',
      size: formatSize(downloads.linuxARM64.size),
      format: getFileFormat(downloads.linuxARM64.link),
      link: downloads.linuxARM64.link,
      recommended: false
    })
  }
}

// 格式化数据
const formatData = (data: DownloadData) => {
  const releases = data[selectedProduct.value]
  if (!releases || releases.length === 0) {
    ElMessage.error('未找到该产品的下载信息')
    return
  }

  allReleaseData = releases

  // 提取所有版本号
  const versions = releases.map(item => item.version)
  availableVersions.value = Array.from(new Set(versions)).sort((a, b) => {
    const aParts = a.split('.').map(Number)
    const bParts = b.split('.').map(Number)
    for (let i = 0; i < Math.max(aParts.length, bParts.length); i++) {
      const aNum = aParts[i] || 0
      const bNum = bParts[i] || 0
      if (aNum !== bNum) {
        return bNum - aNum
      }
    }
    return 0
  })

  // 设置默认选中最新版本
  selectedVersion.value = availableVersions.value[0]
  setDownloadData()
}

// 加载版本信息
const loadVersions = async () => {
  loadingVersions.value = true
  showVersionDropdown.value = false
  try {
    const res = await getVersion(selectedProduct.value)
    formatData(res)
  } catch (error) {
    ElMessage.error('加载版本信息失败')
    console.error('加载版本信息失败:', error)
  } finally {
    loadingVersions.value = false
  }
}

// 选择版本
const selectVersion = (version: string) => {
  selectedVersion.value = version
  showVersionDropdown.value = false
}


// 选择目录
const selectDir = async () => {
  try {
    const dir = await SelectDirectory()

    if (dir) {
      selectedDir.value = dir
      ElMessage.success(`已选择目录: ${dir}`)
    }
  } catch (error) {
    console.error('选择目录失败:', error)
    ElMessage.error('选择目录失败')
  }
}

// 监听版本变化
watch(selectedVersion, () => {
  if (selectedVersion.value) {
    setDownloadData()
  }
})

// 检测操作系统
onMounted(() => {
  const userAgent = navigator.userAgent.toLowerCase()
  if (userAgent.includes('mac')) {
    detectedOS.value = 'mac'
    selectedOS.value = 'mac'
  } else if (userAgent.includes('linux')) {
    detectedOS.value = 'linux'
    selectedOS.value = 'linux'
  } else {
    detectedOS.value = 'windows'
    selectedOS.value = 'windows'
  }

  // 初始加载数据
  loadVersions()

  // 监听下载进度事件
  EventsOn('download:progress', (data: any) => {
    const item = downloadList.value.find(item => item.filePath === data.filePath)
    if (item) {
      item.downloadProgress = Math.round(data.percent)
      item.status = 'downloading'
    }
  })

  // 监听下载完成事件
  EventsOn('download:complete', (data: any) => {
    const item = downloadList.value.find(item => item.filePath === data.filePath)
    if (item) {
      item.downloadProgress = 100
    }
  })

  // 监听下载错误事件
  EventsOn('download:error', (data: any) => {
    const item = downloadList.value.find(item => item.filePath === data.filePath)
    if (item) {
      item.status = 'pending'
      ElMessage.error(`下载失败: ${data.error}`)
    }
  })

  // 监听安装开始事件
  EventsOn('install:start', (data: any) => {
    const item = downloadList.value.find(item => item.filePath === data.filePath)
    if (item) {
      item.status = 'installing'
      item.installProgress = 0
    }
  })

  // 监听安装进度事件
  EventsOn('install:progress', (data: any) => {
    const item = downloadList.value.find(item => item.filePath === data.filePath)
    if (item) {
      item.installProgress = Math.round(data.percent)
    }
  })

  // 监听安装完成事件
  EventsOn('install:complete', (data: any) => {
    const item = downloadList.value.find(item => item.filePath === data.filePath)
    if (item) {
      item.installProgress = 100
      item.status = 'installed'
      ElMessage.success(`安装完成`)
    }
  })

  // 监听安装错误事件
  EventsOn('install:error', (data: any) => {
    const item = downloadList.value.find(item => item.filePath === data.filePath)
    if (item) {
      item.status = 'pending'
      ElMessage.error(`安装失败: ${data.error}`)
    }
  })
})

// 组件卸载时移除事件监听器
onUnmounted(() => {
  EventsOff('download:progress')
  EventsOff('download:complete')
  EventsOff('download:error')
  EventsOff('install:start')
  EventsOff('install:progress')
  EventsOff('install:complete')
  EventsOff('install:error')
})

// 当前下载选项
const currentDownloadOptions = computed(() => {
  return downloadData[selectedOS.value as keyof typeof downloadData] || []
})

// 下载文件
const downloadFile = async (url: string) => {
  // 如果没有选择目录，提示用户先选择目录
  if (!selectedDir.value) {
    ElMessage.warning('请先选择安装目录')
    await selectDir()
    if (!selectedDir.value) {
      return
    }
  }

  try {
    // 从 URL 中提取文件名
    const urlObj = new URL(url)
    const pathname = urlObj.pathname
    const filename = pathname.substring(pathname.lastIndexOf('/') + 1)

    // 检查下载地址是否能够访问
    const res = await CheckDownloadURL(urlObj.href)

    if (res !== null){
      ElMessage.error(`下载地址不可访问请官网自行下载: 原因-${res}`)
      return
    }

    // 临时下载路径（根据操作系统使用正确的路径分隔符）
    const pathSeparator = selectedDir.value.includes('\\') ? '\\' : '/'
    const tempPath = `${selectedDir.value}${pathSeparator}${filename}`

    // 查找对应的下载选项以获取文件大小
    const option = currentDownloadOptions.value.find(opt => opt.link === url)
    const size = option?.size || '未知大小'

    // 检查是否已存在（使用 filePath 检查，避免重复下载到同一位置）
    const exists = downloadList.value.some(item => item.filePath === tempPath)
    if (!exists) {
      // 使用 addToDownloadList 添加到列表
      addToDownloadList(urlObj.href, filename, size, tempPath)
    }

    // 查找列表项并更新状态（使用 filePath 作为唯一标识）
    const item = downloadList.value.find(item => item.filePath === tempPath)
    if (item) {
      item.status = 'downloading'
      item.downloadProgress = 0
    }

    ElMessage.info(`开始下载并安装: ${filename}`)

    // 调用后端的下载并安装函数（使用修改后的 URL）
    DownloadAndInstall(urlObj.href, tempPath, selectedDir.value).catch(error => {
      console.error('下载或安装失败:', error)
    })

  } catch (error) {
    console.error('下载失败:', error)
    ElMessage.error(`下载失败: ${error}`)
  }
}

// 返回首页
const goBack = () => {
  window.history.back()
}
</script>