<template>
  <div class="space-y-4">
    <button
      @click="createSession"
      :disabled="isLoading"
      class="w-full bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 text-white font-bold py-3 px-4 rounded-lg transition-colors"
    >
      {{ isLoading ? 'Создание сессии...' : 'Создать сессию' }}
    </button>
    
    <div v-if="sessionId" class="bg-gray-100 rounded-lg p-4">
      <div class="flex items-center justify-between">
        <div>
          <p class="text-sm text-gray-600 mb-1">ID сессии:</p>
          <p class="font-mono text-lg font-bold text-gray-800">{{ sessionId }}</p>
        </div>
        <button
          @click="copySessionId"
          class="bg-gray-200 hover:bg-gray-300 text-gray-700 px-3 py-2 rounded-lg transition-colors"
        >
          Копировать
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { ApiService } from '@/services/api'

const isLoading = ref(false)
const sessionId = ref<string | null>(null)

const createSession = async () => {
  isLoading.value = true
  try {
    const response = await ApiService.createSession()
    sessionId.value = response.id
  } catch (error) {
    console.error('Failed to create session:', error)
    alert('Ошибка при создании сессии')
  } finally {
    isLoading.value = false
  }
}

const copySessionId = async () => {
  if (sessionId.value) {
    try {
      await navigator.clipboard.writeText(sessionId.value)
      alert('ID сессии скопирован!')
    } catch (error) {
      console.error('Failed to copy session ID:', error)
    }
  }
}
</script> 