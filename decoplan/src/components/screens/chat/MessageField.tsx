"use client"
import { ArrowRightToLine, Send } from "lucide-react"
import { useState } from "react"
import Field from "../../ui/field/Field"

interface IMessageField {}

export function MessageField({}: IMessageField) {
	const [message, setMessage] = useState<string>("")

	const onSubmit = () => {
		message && setMessage("Хуй")
	}
	return (
		<div className='border-t border-border p-layout flex items-center justify-between'>
			<Field
				className='w-4/5 bg-red-500'
				placeholder='Write a message...'
				Icon={ArrowRightToLine}
				value={message}
				onChange={(e) => setMessage(e.target.value)}
			/>
			<button
				className='hover:text-primary transition-colors'
				onClick={onSubmit}
				disabled={!message}
			>
				<Send />
			</button>
		</div>
	)
}
