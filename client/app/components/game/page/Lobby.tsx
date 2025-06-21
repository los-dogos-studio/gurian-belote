import { useState } from 'react';
import { useGameClient } from '../GameClientContext'
import Panel from '~/components/Panel';
import InputField from '~/components/InputField';
import Button from '~/components/Button';
import { LuCrown, LuLogIn } from 'react-icons/lu';
import Break from '~/components/Break';

const Lobby = () => {
	const [userId, setUserId] = useState('');
	const [roomId, setRoomId] = useState('');
	const client = useGameClient();

	const handleJoinRoom = () => {
		client.connect(userId).then(() => {
			client.joinRoom(roomId);
		});
	};

	const handleCreateRoom = () => {
		client.connect(userId).then(() => {
			client.createRoom();
		});
	};

	const Title = () => {
		return (
			<h1 className="text-4xl font-bold text-amber-400/90 text-center mb-8 tracking-wider">
				Gurian Belote
			</h1>
		);
	}

	return (
		<div className="w-full h-full grid place-items-center">
			<Panel className="max-w-md w-full border border-yellow-600/30 rounded-2xl p-8">
				<Title />
				<InputField
					label={'Player Name'}
					placeholder={'Choose your alias...'}
					value={userId}
					onChange={setUserId}
				/>
				<Button
					onClick={handleCreateRoom}
					disabled={!userId}
					className="mb-4"
				>
					<LuCrown className="w-5 h-5 mr-2" />
					Create Room
				</Button>

				<Break label={'OR'} />

				<InputField
					label={'Invite Code'}
					placeholder={'Enter invitation code...'}
					value={roomId}
					onChange={setRoomId}
					disabled={!userId}
				/>
				<Button
					onClick={handleJoinRoom}
					disabled={!userId || !roomId}
					variant="secondary"
					className="mb-4"
				>
					<LuLogIn className="w-5 h-5 mr-2" />
					Join Room
				</Button>
			</Panel>
		</div>
	);
};

export default Lobby;

