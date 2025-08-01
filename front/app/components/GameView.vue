<template>
  <div class="game-container" v-if="game">
    <div class="guesses">
      <div>
        <h3>{{ myName }}</h3>
        <ul>
          <li v-for="(guess, i) in myGuesses" :key="i">{{ guess }}</li>
        </ul>
      </div>

      <div>
        <h3>{{ opponentName }}</h3>
        <ul>
          <li v-for="(guess, i) in opponentGuesses" :key="i">{{ guess }}</li>
        </ul>
      </div>
    </div>

    <div v-if="game.status === 'NEW'" class="input-area">
      <input
        v-model="guessInput"
        @keyup.enter="submitGuess"
        placeholder="Введите слово"
      />
      <button @click="submitGuess">Отправить</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue';
import type {GameState} from "~/composables/useGameSocket";

const props = defineProps<{
  api: {
    game: Ref<GameState | null>;
    sendGuess: (word: string) => void;
  };
}>();

const guessInput = ref('');
const game = props.api.game;

const myId = computed(() => game.value?.you || '');
const otherId = computed(() =>
  Object.keys(game.value?.players || {}).find((id) => id !== myId.value) || ''
);

const myName = computed(() => game.value?.players?.[myId.value]?.username || '');
const opponentName = computed(() => game.value?.players?.[otherId.value]?.username || '');

const myGuesses = computed(() => game.value?.guesses?.[myId.value] || []);
const opponentGuesses = computed(() => game.value?.guesses?.[otherId.value] || []);

const submitGuess = () => {
  if (!guessInput.value.trim()) return;
  props.api.sendGuess(guessInput.value.trim());
  guessInput.value = '';
};
</script>

<style scoped>
.game-container {
  max-width: 800px;
  margin: 0 auto;
  text-align: center;
}
.guesses {
  display: flex;
  justify-content: space-around;
  margin-bottom: 20px;
}
.input-area {
  margin-top: 20px;
}
input {
  margin-bottom: 10px;
  padding: 8px;
}
</style>
