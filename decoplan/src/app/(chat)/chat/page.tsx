import { CreateChat } from "@/components/ui/createChat/CreateChat"

export default function ChatsPage() {
	return (
		<>
			<p className='p-layout'> Click chat on the left side for open</p>
			<div className='fixed bottom-6 right-6'>
				<CreateChat />
			</div>
		</>
	)
}
