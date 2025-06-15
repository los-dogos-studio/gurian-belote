import { useState } from 'react';
import { useGameClient } from '../GameClientContext'

const Lobby = () => {
	const [userId, setUserId] = useState('');
	const [roomId, setRoomId] = useState('');
	const client = useGameClient();

	const handleJoinRoom = () => {
		console.log('Joining room:', roomId, 'as user:', userId);
		client.connect(userId).then(() => {
			client.joinRoom(roomId);
		});
	};

	const handleCreateRoom = () => {
		console.log('Creating room as user:', userId);
		client.connect(userId).then(() => {
			client.createRoom();
		});
	};

	return (
		<div className="flex items-center justify-center min-h-screen bg-gradient-to-br from-[#1a1a1a] via-[#2c0a0a] to-[#1a1a1a] text-gold-100 font-serif">
			<div className="bg-[#1e0f0f] p-10 rounded-xl shadow-2xl border-2 border-gold-600 w-full max-w-md">
				<h1 className="text-3xl font-extrabold text-center mb-8 text-gold-400 drop-shadow-glow animate-pulse">
					♠️ Gurian Belote ♣️
				</h1>

				<div className="mb-6">
					<label className="block text-sm font-medium text-gold-200 mb-1">🎲 User ID</label>
					<input
						type="text"
						className="w-full px-4 py-2 bg-[#2c1a1a] border border-gold-500 rounded-lg text-white placeholder:text-gold-300 focus:outline-none focus:ring-2 focus:ring-gold-400"
						placeholder="Enter your player name"
						value={userId}
						onChange={(e) => setUserId(e.target.value)}
					/>
				</div>

				<div className="mb-6">
					<label className="block text-sm font-medium text-gold-200 mb-1">🃏 Room ID</label>
					<input
						type="text"
						className="w-full px-4 py-2 bg-[#2c1a1a] border border-gold-500 rounded-lg text-white placeholder:text-gold-300 focus:outline-none focus:ring-2 focus:ring-gold-400"
						placeholder="Enter room ID"
						value={roomId}
						onChange={(e) => setRoomId(e.target.value)}
					/>
				</div>

				<div className="flex flex-col sm:flex-row gap-4">
					<button
						onClick={handleJoinRoom}
						className="w-full bg-gold-600 hover:bg-gold-500 text-black font-bold py-2 px-4 rounded-lg shadow-md transition hover:scale-105"
					>
						🎯 Join Room
					</button>
					<button
						onClick={handleCreateRoom}
						className="w-full bg-green-500 hover:bg-green-400 text-black font-bold py-2 px-4 rounded-lg shadow-md transition hover:scale-105"
					>
						🏗️ Create Room
					</button>
				</div>
			</div>
		</div>
	);
};

export default Lobby;

