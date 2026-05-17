import type { Session, SessionManager } from '@gurian-belote/lib'

const STORAGE_KEY = 'game-client-session'

export class WebSessionManager implements SessionManager {
	saveSession(sessionId: string, userId: string): Promise<void> {
		const session: Session = { id: sessionId, userId }
		localStorage.setItem(STORAGE_KEY, JSON.stringify(session))
		return Promise.resolve()
	}

	getSession(): Promise<Session | null> {
		const raw = localStorage.getItem(STORAGE_KEY)
		return Promise.resolve(raw ? (JSON.parse(raw) as Session) : null)
	}

	clearSession(): Promise<void> {
		localStorage.removeItem(STORAGE_KEY)
		return Promise.resolve()
	}
}
