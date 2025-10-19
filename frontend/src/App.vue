<script lang="ts" setup>
import {useDrvice} from "./store/useDrvice";
import {Actions, Clean} from "../wailsjs/go/main/App";
import { ElMessage } from 'element-plus'
import {ref} from "vue";

const drviceStore = useDrvice();
drviceStore.getDrviceInfo()

const data = ref()

const actions = async () => {
  data.value = await Actions();
}

const clean = async () => {
 const data = await Clean();
  console.log(data);
  ElMessage(data)
}
</script>

<template>
{{data}}
  <div>
    <button @click="actions">actions</button>
  </div>
  <button @click="clean">clean</button>
  <router-view />
</template>

<style>
</style>
