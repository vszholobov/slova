import { useRuntimeConfig } from '#app';

let socket: WebSocket | null = null;

export default function useSession() {
  const config = useRuntimeConfig();

  const connectToSession = (sessionId: string, username: string): Promise<void> => {
    return new Promise((resolve, reject) => {
      const wsUrl = `${config.public.apiWsSchema}${config.public.apiBaseUrl}/session/${sessionId}`;
      socket = new WebSocket(wsUrl);

      socket.onopen = () => {
        console.log('Подключено к веб-сокету');
        const message = {
          command: 'SET_USERNAME',
          body: { username },
        };
        socket?.send(JSON.stringify(message));
        resolve();
      };

      socket.onerror = (event) => {
        console.error('Ошибка веб-сокета:', event);
        reject(event);
      };

      socket.onclose = () => {
        console.log('Соединение с веб-сокетом закрыто');
        socket = null;
      };
    });
  };

  return {
    connectToSession,
  };
}
