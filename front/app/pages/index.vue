<template>
  <div class="container">
    <h1>Добро пожаловать в игру в слова!</h1>

    <div v-if="!connected">
      <button @click="createSession" :disabled="loadingCreate">
        {{ loadingCreate ? 'Создание...' : 'Создать сессию' }}
      </button>

      <div v-if="sessionId">
        <p>Сессия создана! ID: <strong>{{ sessionId }}</strong></p>
      </div>

      <button @click="showModal = true">Подключиться</button>

      <ConnectSessionModal
        v-if="showModal"
        @connected="onConnected"
        @close="showModal = false"
      />
    </div>

    <GameView v-if="connected && showGameView" :api="gameApi" />
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';
import GameView from '~/components/GameView.vue';
import ConnectSessionModal from '~/components/ConnectSessionModal.vue';
import { useGameSocket } from '~/composables/useGameSocket';

const sessionId = ref<string | null>(null);
const connected = ref(false);
const showModal = ref(false);
const loadingCreate = ref(false);
const gameApi = ref<ReturnType<typeof useGameSocket> | null>(null);

const showGameView = computed(() => {
  const api = gameApi.value;
  if (!api || !api.game || !api.game.value) return false;
  const status = api.game.value.status;
  return status === 'NEW' || status === 'ACTIVE';
});


const createSession = async () => {
  loadingCreate.value = true;
  try {
    const data = await $fetch<{ id: string }>('/api/session', {
      method: 'POST',
    });
    sessionId.value = data.id;
  } catch (e) {
    alert('Не удалось создать сессию');
    console.error(e);
  } finally {
    loadingCreate.value = false;
  }
};

const onConnected = (ws: WebSocket) => {
  connected.value = true;
  showModal.value = false;
  gameApi.value = useGameSocket(ws);
};
</script>

<style scoped>
.container {
  max-width: 600px;
  margin: 0 auto;
  text-align: center;
}
button {
  margin: 10px;
  padding: 8px 16px;
}
</style>
