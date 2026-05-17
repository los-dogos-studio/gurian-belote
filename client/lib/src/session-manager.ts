export interface Session {
	id: string
	userId: string
}

export interface SessionManager {
	saveSession(sessionId: string, userId: string): Promise<void>
	getSession(): Promise<Session | null>
	clearSession(): Promise<void>
}
