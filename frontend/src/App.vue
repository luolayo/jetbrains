<script lang="ts" setup>
import {useDrvice} from "./store/useDrvice";
import {Actions, Clean} from "../wailsjs/go/main/App";
import { ElMessage } from 'element-plus'

const drviceStore = useDrvice();
drviceStore.getDrviceInfo()

const actions = async () => {
 const data = await Actions();
 
  console.log(data);
}

const clean = async () => {
 const data = await Clean();
  console.log(data);
  ElMessage(data)
}
</script>

<template>
  <div>
    <button @click="actions">actions</button>
  </div>
  <button @click="clean">clean</button>
  <router-view />
</template>

<style>
</style>
