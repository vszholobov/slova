<template>
  <div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
    <div class="bg-white rounded-lg p-6 w-full max-w-md mx-4">
      <h2 class="text-xl font-bold mb-4">Подключиться к сессии</h2>
      
      <div class="space-y-4">
        <div>
          <label for="sessionId" class="block text-sm font-medium text-gray-700 mb-2">
            ID сессии
          </label>
          <input
            id="sessionId"
            v-model="sessionId"
            type="text"
            placeholder="Введите ID сессии"
            class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
            @keyup.enter="connect"
          />
        </div>
        
        <div class="flex space-x-3">
          <button
            @click="connect"
            :disabled="!sessionId.trim()"
            class="flex-1 bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 text-white font-bold py-2 px-4 rounded-lg transition-colors"
          >
            Подключиться
          </button>
          <button
            @click="$emit('close')"
            class="flex-1 bg-gray-300 hover:bg-gray-400 text-gray-700 font-bold py-2 px-4 rounded-lg transition-colors"
          >
            Отмена
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const emit = defineEmits<{
  close: []
  connect: [sessionId: string]
}>()

const sessionId = ref('')

const connect = () => {
  if (sessionId.value.trim()) {
    emit('connect', sessionId.value.trim())
  }
}
</script>