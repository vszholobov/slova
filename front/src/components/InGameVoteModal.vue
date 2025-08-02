<template>
  <div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
    <div class="bg-white rounded-lg p-6 w-full max-w-md mx-4">
      <h2 class="text-xl font-bold mb-4 text-center">
        {{ voteTitle }}
      </h2>
      
      <div class="space-y-4">
        <div class="flex justify-between items-center">
          <span class="font-medium">Игрок 1:</span>
          <div class="flex items-center space-x-2">
            <span v-if="gameState.vote?.player1 === true" class="text-green-600 text-xl">✅</span>
            <span v-else-if="gameState.vote?.player1 === false" class="text-red-600 text-xl">❌</span>
            <span v-else class="text-gray-400">⏳</span>
          </div>
        </div>
        
        <div class="flex justify-between items-center">
          <span class="font-medium">Игрок 2:</span>
          <div class="flex items-center space-x-2">
            <span v-if="gameState.vote?.player2 === true" class="text-green-600 text-xl">✅</span>
            <span v-else-if="gameState.vote?.player2 === false" class="text-red-600 text-xl">❌</span>
            <span v-else class="text-gray-400">⏳</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useGameStore } from '@/store/game'

const gameStore = useGameStore()
const gameState = computed(() => gameStore.gameState!)

const voteTitle = computed(() => {
  if (!gameState.value.vote) return ''
  switch (gameState.value.vote.type) {
    case 'STOP_WIN':
      return 'Признать победу'
    case 'STOP_SKIP':
      return 'Пропустить игру'
    default:
      return 'Голосование'
  }
})
</script>