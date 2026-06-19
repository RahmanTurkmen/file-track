import { defineStore } from 'pinia'
import api from '../services/api'

export const useFilesStore = defineStore('files', {
  state: () => ({
    files: [],
    search: '',
    activeTab: 'mydrive'
  }),

  getters: {
    filteredFiles(state) {
      const files = Array.isArray(state.files) ? state.files : []

      let result = files.filter(f =>
        (f.name || '').toLowerCase().includes(state.search.toLowerCase())
      )

      if (state.activeTab === 'mydrive') {
        return result.filter(f => f.status !== 'deleted')
      }

      if (state.activeTab === 'trash') {
        return result.filter(f => f.status === 'deleted')
      }

      if (state.activeTab === 'recent') {
        return result
          .filter(f => f.status !== 'deleted')
          .slice()
          .sort((a, b) => b.id - a.id)
      }

      return result
    }
  },

  actions: {

    async fetchFiles() {
      const res = await api.get('/files')
      this.files = Array.isArray(res.data) ? res.data : []
    },

    async uploadFile(name, size) {
      await api.post('/files', { name, size })
      await this.fetchFiles()
    },

    async deleteFile(id) {
      await api.delete(`/files/${id}`)

      const res = await api.get('/files')
      this.files = res.data
    }
  }
})