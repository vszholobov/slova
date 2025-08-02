import type { CreateSessionResponse } from '@/types'

const API_BASE_URL = 'http://localhost:80'

export class ApiService {
  static async createSession(): Promise<CreateSessionResponse> {
    const response = await fetch(`${API_BASE_URL}/session/`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
    })

    if (!response.ok) {
      throw new Error('Failed to create session')
    }

    return response.json()
  }
} 