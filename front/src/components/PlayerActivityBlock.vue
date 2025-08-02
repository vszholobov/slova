<template>
  <div class="bg-gray-50 rounded-lg p-4">
    <div class="flex items-center justify-between mb-4">
      <h3 class="text-lg font-semibold text-gray-800">
        {{ player.username }}
      </h3>
      <span
        v-if="isCurrentPlayer"
        class="bg-blue-100 text-blue-800 text-xs font-medium px-2 py-1 rounded-full"
      >
        Вы
      </span>
    </div>
    
    <!-- История слов -->
    <div class="mb-4">
      <h4 class="text-sm font-medium text-gray-600 mb-2">История слов:</h4>
      <div class="space-y-1">
        <div
          v-for="(guess, index) in guesses"
          :key="index"
          class="bg-white px-3 py-2 rounded border text-sm"
        >
          {{ guess }}
        </div>
        <div v-if="guesses.length === 0" class="text-gray-400 text-sm italic">
          Пока нет слов
        </div>
      </div>
    </div>
    
    <!-- Текущая догадка -->
    <CurrentGuess
      :is-current-player="isCurrentPlayer"
      :is-active="isActive"
      :last-guess="guesses[guesses.length - 1]"
    />
  </div>
</template>

<script setup lang="ts">
interface Props {
  player: { username: string }
  guesses: string[]
  isCurrentPlayer: boolean
  isActive: boolean
}

defineProps<Props>()
</script>