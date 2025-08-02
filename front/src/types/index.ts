export interface GameState {
  you: "1" | "2";
  status: "NEW" | "ACTIVE" | "WIN" | "SKIP";
  vote: {
    type: "STOP_WIN" | "STOP_SKIP";
    player1: boolean | null;
    player2: boolean | null;
  } | null;
  guesses: {
    player1: string[];
    player2: string[];
  };
  players: {
    player1: { username: string };
    player2: { username: string };
  };
  history: {
    status: "WIN" | "SKIP";
    guesses: {
      player1: string[];
      player2: string[];
    };
  }[];
}

export interface CreateSessionResponse {
  id: string;
}

export interface UserCommand {
  command: string;
  body: any;
}

export interface GuessCommandBody {
  guess: string;
}

export interface VoteCommandBody {
  vote: boolean;
}

export interface SetUsernameCommandBody {
  username: string;
} 