import { ref } from 'vue';

export interface GameState {
  you: string;
  status: 'WIN' | 'SKIP' | 'ACTIVE' | 'NEW';
  guesses: Record<string, string[]>;
  players: Record<string, { username: string }>;
}

const game = ref<GameState | null>(null);

export function useGameSocket(ws: WebSocket) {
  const sendGuess = (word: string) => {
    if (ws.readyState === WebSocket.OPEN) {
      ws.send(
        JSON.stringify({
          command: 'MAKE_GUESS',
          body: { word },
        })
      );
    }
  };

  ws.addEventListener('message', (event) => {
    try {
      const message = JSON.parse(event.data);
      if (message.type === 'GAME_STATE') {
        game.value = message.body;
      }
    } catch (e) {
      console.error('Ошибка парсинга WebSocket:', e);
    }
  });

  return {
    game,
    sendGuess,
  };
}
