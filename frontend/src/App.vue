<script setup>
import { onMounted, ref } from 'vue'
import { useFilesStore } from './stores/files'

const store = useFilesStore()

const name = ref("")
const size = ref(0)

onMounted(() => {
  store.fetchFiles()

  setInterval(() => {
    store.fetchFiles()
  }, 2000)
})

function upload() {
  if (name.value.trim()) {
    store.uploadFile(name.value, size.value || 10)
    name.value = ""
    size.value = 0
  }
}
</script>

<template>
  <div class="app">

    <div class="sidebar">

      <h2>STORAGE</h2>

      <div class="sidebar-item" @click="store.activeTab='mydrive'">
        My Drive
      </div>

      <div class="sidebar-item" @click="store.activeTab='recent'">
        Recent
      </div>

      <div class="sidebar-item" @click="store.activeTab='trash'">
        Trash
      </div>

    </div>

    <div class="main">

      <div class="header">
        <div class="title">File Tracker System</div>
      </div>

      <div style="display:flex; gap:10px; margin-bottom:20px;">

        <input v-model="name" placeholder="File name" />
        <input v-model="size" type="number" placeholder="Size MB" style="width:120px"/>

        <button class="btn-primary" @click="upload">
          Upload
        </button>

      </div>

      <input v-model="store.search" placeholder="Search files..." />

      <div class="card">

        <table>

          <thead>
            <tr>
              <th>Name</th>
              <th>Size</th>
              <th>Status</th>
              <th>Date</th>
              <th>Action</th>
            </tr>
          </thead>

          <tbody>
            <tr v-for="f in store.filteredFiles" :key="f.id">

              <td> {{ f.name }}</td>
              <td>{{ f.size }} MB</td>

              <td>
                <span :class="['badge', f.status]">
                  {{ f.status }}
                </span>
              </td>

              <td>{{ f.created_at }}</td>

              <td>
                <button class="btn-danger"
                        @click="store.deleteFile(f.id)">
                  delete
                </button>
              </td>

            </tr>
          </tbody>

        </table>

      </div>

    </div>

  </div>
</template>