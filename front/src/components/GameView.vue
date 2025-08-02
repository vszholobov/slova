<template>
  <div class="max-w-4xl mx-auto">
    <div class="bg-white rounded-lg shadow-xl p-6">
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- Игрок 2 (оппонент) -->
        <PlayerActivityBlock
          :player="gameState.players.player2"
          :guesses="gameState.guesses.player2"
          :is-current-player="false"
          :is-active="isGameActive"
        />
        
        <!-- Игрок 1 (текущий) -->
        <PlayerActivityBlock
          :player="gameState.players.player1"
          :guesses="gameState.guesses.player1"
          :is-current-player="true"
          :is-active="isGameActive"
        />
      </div>
      
      <!-- Меню игры -->
      <GameMenu />
      
      <!-- Модальное окно голосования -->
      <InGameVoteModal v-if="hasVote" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useGameStore } from '@/store/game'
import PlayerActivityBlock from './PlayerActivityBlock.vue'
import GameMenu from './GameMenu.vue'
import InGameVoteModal from './InGameVoteModal.vue'

const gameStore = useGameStore()

const gameState = computed(() => gameStore.gameState!)
const isGameActive = computed(() => gameStore.isGameActive)
const hasVote = computed(() => gameStore.hasVote)
</script>