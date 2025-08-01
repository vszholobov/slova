<template>
  <button @click="createSession" :disabled="loading">
    {{ loading ? 'Создание...' : 'Создать сессию' }}
  </button>
</template>

<script setup lang="ts">
import { ref, defineEmits } from 'vue';
import { useFetch } from '#imports'; // fixme: Component is already mounted, please use $fetch instead. https://nuxt.com/docs/getting-started/data-fetching

interface CreateSessionResponse {
  id: string;
}

const loading = ref(false);
const emit = defineEmits<{ (e: 'sessionCreated', id: string): void }>();

const createSession = async () => {
  loading.value = true;
  try {
    const { data } = await useFetch<CreateSessionResponse>('api/session', {
      method: 'POST',
    });
    if (data.value?.id) {
      emit('sessionCreated', data.value.id);
    }
  } catch (error) {
    console.error('Ошибка при создании сессии:', error);
    alert('Не удалось создать сессию');
  } finally {
    loading.value = false;
  }
};
</script>
