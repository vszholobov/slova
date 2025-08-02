import type { GameState, UserCommand } from '@/types'

type GameStateHandler = (state: GameState) => void

export class WebSocketService {
  private ws: WebSocket | null = null
  private onGameStateUpdate: GameStateHandler | null = null

  connect(sessionId: string, onStateUpdate: GameStateHandler) {
    this.onGameStateUpdate = onStateUpdate

    // Открываем обычный WebSocket
    this.ws = new WebSocket(`ws://localhost:80/session/${sessionId}`)

    this.ws.onopen = () => {
      console.log('WebSocket connected')
    }

    this.ws.onclose = () => {
      console.log('WebSocket disconnected')
    }

    this.ws.onerror = (e) => {
      console.error('WebSocket error', e)
    }

    this.ws.onmessage = (event) => {
      try {
        const data = JSON.parse(event.data)
        if (this.onGameStateUpdate) {
          this.onGameStateUpdate(data)
        }
      } catch (err) {
        console.error('Failed to parse WS message', err)
      }
    }
  }

  disconnect() {
    if (this.ws) {
      this.ws.close()
      this.ws = null
    }
  }

  sendCommand(command: UserCommand) {
    if (this.ws && this.ws.readyState === WebSocket.OPEN) {
      this.ws.send(JSON.stringify(command))
    }
  }

  sendGuess(guess: string) {
    this.sendCommand({
      command: 'GUESS',
      body: { guess }
    })
  }

  sendVote(vote: boolean, command: 'START_NEW' | 'STOP_WIN' | 'STOP_SKIP') {
    this.sendCommand({
      command,
      body: { vote }
    })
  }

  sendUsername(username: string) {
    this.sendCommand({
      command: 'SET_USERNAME',
      body: { username }
    })
  }
} 