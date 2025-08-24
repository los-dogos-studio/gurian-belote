interface StoredSession {
	roomId: string;
	userName: string;
	autoReconnect: boolean;
}

export class SessionManager {
	private readonly storageKey = 'session_info';

	public getSession(): StoredSession | null {
		try {
			const stored = localStorage.getItem(this.storageKey);
			if (!stored) return null;
			return JSON.parse(stored) as StoredSession;
		} catch {
			return null;
		}
	}

	public saveSession(session: Partial<StoredSession>): void {
		const current = this.getSession() || { roomId: '', userName: '', autoReconnect: true };
		const updated = { ...current, ...session };
		localStorage.setItem(this.storageKey, JSON.stringify(updated));
	}

	public clearSession(): void {
		localStorage.removeItem(this.storageKey);
	}

	public setAutoReconnect(enabled: boolean): void {
		const session = this.getSession();
		if (session) {
			session.autoReconnect = enabled;
			this.saveSession(session);
		}
	}
}