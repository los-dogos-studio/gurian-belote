interface Session {
	id: string
	userId: string
}

export class SessionManager {
	private readonly storageKey = 'game-client-session'

	public saveSession(sessionId: string, userId: string): void {
		const session = { id: sessionId, userId: userId }
		localStorage.setItem(this.storageKey, JSON.stringify(session))
	}

	public getSession(): Session | null {
		const sessionJson = localStorage.getItem(this.storageKey)
		return sessionJson ? (JSON.parse(sessionJson) as Session) : null
	}

	public clearSession(): void {
		localStorage.removeItem(this.storageKey)
	}
}
