<template>
  <div>
    <button
      @click="showModal = true"
      class="w-full bg-green-600 hover:bg-green-700 text-white font-bold py-3 px-4 rounded-lg transition-colors"
    >
      Подключиться к сессии
    </button>
    
    <ConnectToSessionModal
      v-if="showModal"
      @close="showModal = false"
      @connect="handleConnect"
    />
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import ConnectToSessionModal from './ConnectToSessionModal.vue'
import { useGameStore } from '@/store/game'
import { WebSocketService } from '@/services/websocket'

const showModal = ref(false)
const gameStore = useGameStore()
const wsService = new WebSocketService()

const handleConnect = (sessionId: string) => {
  gameStore.setSessionId(sessionId)
  
  wsService.connect(sessionId, (gameState) => {
    gameStore.setGameState(gameState)
    gameStore.setConnectionStatus(true)
  })
  
  showModal.value = false
}
</script>