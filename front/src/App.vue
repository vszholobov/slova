<template>
  <div id="app" class="min-h-screen bg-gradient-to-br from-blue-500 to-purple-600">
    <div class="container mx-auto px-4 py-8">
      <component :is="currentComponent" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useGameStore } from '@/store/game'
import MainMenu from '@/components/MainMenu.vue'
import GameView from '@/components/GameView.vue'

const gameStore = useGameStore()

const currentComponent = computed(() => {
  if (gameStore.isConnected && gameStore.gameState) {
    return GameView
  }
  return MainMenu
})
</script> 