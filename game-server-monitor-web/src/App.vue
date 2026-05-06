<template>
  <div class="page">
    <div class="header">
      <div>
        <h1>Game Server Monitor</h1>
        <p>轻量级游戏服务器监控后台</p>
      </div>
      <el-button type="primary" @click="loadAll">刷新</el-button>
    </div>

    <el-row :gutter="16" class="cards">
      <el-col :span="6"><el-card><div class="stat-title">服务器数量</div><div class="stat-value">{{ servers.length }}</div></el-card></el-col>
      <el-col :span="6"><el-card><div class="stat-title">总在线</div><div class="stat-value">{{ totalOnline }}</div></el-card></el-col>
      <el-col :span="6"><el-card><div class="stat-title">告警数</div><div class="stat-value">{{ alerts.length }}</div></el-card></el-col>
      <el-col :span="6"><el-card><div class="stat-title">异常服务器</div><div class="stat-value">{{ abnormalCount }}</div></el-card></el-col>
    </el-row>

    <el-card class="panel">
      <template #header><div class="panel-title">服务器列表</div></template>
      <el-table :data="servers" border style="width: 100%">
        <el-table-column prop="server_id" label="服务器ID" width="120" />
        <el-table-column prop="name" label="名称" width="120" />
        <el-table-column prop="ip" label="IP" width="140" />
        <el-table-column prop="online" label="在线人数" width="110" />
        <el-table-column prop="cpu" label="CPU" width="100">
          <template #default="scope">{{ Number(scope.row.cpu).toFixed(2) }}%</template>
        </el-table-column>
        <el-table-column prop="memory" label="内存" width="100">
          <template #default="scope">{{ Number(scope.row.memory).toFixed(2) }}%</template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="120">
          <template #default="scope">
            <el-tag :type="statusType(scope.row.status)">{{ statusText(scope.row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="last_report" label="最后上报" />
        <el-table-column label="操作" width="220">
          <template #default="scope">
            <el-button size="small" @click="restart(scope.row.server_id)">重启</el-button>
            <el-button size="small" type="warning" @click="maintenance(scope.row.server_id)">维护</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-row :gutter="16">
      <el-col :span="12">
        <el-card class="panel">
          <template #header><div class="panel-title">告警列表</div></template>
          <el-table :data="alerts" border height="300">
            <el-table-column prop="server_id" label="服务器" width="100" />
            <el-table-column prop="type" label="类型" width="100" />
            <el-table-column prop="level" label="等级" width="100" />
            <el-table-column prop="message" label="信息" />
          </el-table>
        </el-card>
      </el-col>

      <el-col :span="12">
        <el-card class="panel">
          <template #header><div class="panel-title">日志搜索</div></template>
          <div class="search-bar">
            <el-input v-model="keyword" placeholder="输入 error / timeout / mysql" clearable />
            <el-button type="primary" @click="searchLogs">查询</el-button>
          </div>
          <div class="log-box">
            <div v-for="(line, index) in logs" :key="index" class="log-line">{{ line }}</div>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { computed, onMounted, ref } from 'vue'
import { ElMessage } from 'element-plus'

const API_BASE = 'http://127.0.0.1:8080/api'

const servers = ref([])
const alerts = ref([])
const logs = ref([])
const keyword = ref('error')

const totalOnline = computed(() => servers.value.reduce((sum, item) => sum + item.online, 0))
const abnormalCount = computed(() => servers.value.filter(item => item.status !== 'running').length)

function statusType(status) {
  if (status === 'running') return 'success'
  if (status === 'warning') return 'warning'
  if (status === 'maintenance') return 'info'
  return 'danger'
}

function statusText(status) {
  const map = { running: '正常', warning: '告警', down: '异常', maintenance: '维护中' }
  return map[status] || status
}

async function request(url, options = {}) {
  const res = await fetch(url, options)
  if (!res.ok) throw new Error(`请求失败: ${res.status}`)
  return res.json()
}

async function loadServers() {
  const res = await request(`${API_BASE}/servers`)
  servers.value = res.data || []
}

async function loadAlerts() {
  const res = await request(`${API_BASE}/alerts`)
  alerts.value = res.data || []
}

async function searchLogs() {
  const res = await request(`${API_BASE}/logs?keyword=${encodeURIComponent(keyword.value)}`)
  logs.value = res.data || []
}

async function restart(serverId) {
  await request(`${API_BASE}/servers/${serverId}/restart`, { method: 'POST' })
  ElMessage.success(`服务器 ${serverId} 重启操作已提交`)
  await loadAll()
}

async function maintenance(serverId) {
  await request(`${API_BASE}/servers/${serverId}/maintenance`, { method: 'POST' })
  ElMessage.success(`服务器 ${serverId} 已切换维护状态`)
  await loadAll()
}

async function loadAll() {
  await Promise.all([loadServers(), loadAlerts(), searchLogs()])
}

onMounted(loadAll)
</script>

<style scoped>
.page { min-height: 100vh; background: #f5f7fb; padding: 24px; box-sizing: border-box; }
.header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 24px; }
.header h1 { margin: 0; font-size: 28px; }
.header p { margin: 8px 0 0; color: #666; }
.cards { margin-bottom: 16px; }
.stat-title { color: #666; font-size: 14px; }
.stat-value { margin-top: 8px; font-size: 28px; font-weight: bold; }
.panel { margin-bottom: 16px; }
.panel-title { font-weight: bold; }
.search-bar { display: flex; gap: 8px; margin-bottom: 12px; }
.log-box { height: 240px; overflow: auto; background: #111827; color: #e5e7eb; padding: 12px; border-radius: 6px; font-family: Consolas, Monaco, monospace; font-size: 13px; }
.log-line { line-height: 1.8; }
</style>
