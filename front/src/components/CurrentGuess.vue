<template>
  <div>
    <h4 class="text-sm font-medium text-gray-600 mb-2">
      {{ isCurrentPlayer ? 'Ваша догадка:' : 'Догадка игрока:' }}
    </h4>
    
    <div v-if="isCurrentPlayer && isActive">
      <input
        v-model="currentGuess"
        @keyup.enter="sendGuess"
        type="text"
        placeholder="Введите слово..."
        :disabled="isBlocked"
        class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 disabled:bg-gray-100"
      />
      <button
        @click="sendGuess"
        :disabled="!currentGuess.trim() || isBlocked"
        class="mt-2 w-full bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 text-white font-bold py-2 px-4 rounded-lg transition-colors"
      >
        Отправить
      </button>
    </div>
    
    <div v-else-if="lastGuess" class="bg-white px-3 py-2 rounded border">
      {{ lastGuess }}
    </div>
    
    <div v-else class="text-gray-400 text-sm italic">
      {{ isCurrentPlayer ? 'Введите слово' : 'Ожидание...' }}
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useGameStore } from '@/store/game'
import { WebSocketService } from '@/services/websocket'

interface Props {
  isCurrentPlayer: boolean
  isActive: boolean
  lastGuess?: string
}

const props = defineProps<Props>()

const gameStore = useGameStore()
const wsService = new WebSocketService()
const currentGuess = ref('')

const isBlocked = computed(() => {
  if (!gameStore.gameState) return true
  
  const { guesses } = gameStore.gameState
  const currentPlayer = gameStore.currentPlayer
  
  if (currentPlayer === '1') {
    return guesses.player1.length > guesses.player2.length
  } else {
    return guesses.player2.length > guesses.player1.length
  }
})

const sendGuess = () => {
  if (currentGuess.value.trim() && !isBlocked.value) {
    wsService.sendGuess(currentGuess.value.trim())
    currentGuess.value = ''
  }
}
</script>