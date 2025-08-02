<template>
  <button
    @click="startNewGame"
    class="bg-green-600 hover:bg-green-700 text-white font-bold py-3 px-6 rounded-lg transition-colors"
  >
    Начать новую игру
  </button>
</template>

<script setup lang="ts">
import { WebSocketService } from '@/services/websocket'

const wsService = new WebSocketService()

const startNewGame = () => {
  wsService.sendVote(true, 'START_NEW')
}
</script> 