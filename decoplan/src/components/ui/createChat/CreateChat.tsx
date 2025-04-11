import { Pencil } from "lucide-react"

export function CreateChat() {
	return (
		<button
			className='bg-indigo-500 text-white rounded-full p-4 shadow-lg hover:bg-indigo-600'
			aria-label='Create new chat'
		>
			<Pencil className='w-5 h-5' />
		</button>
	)
}
