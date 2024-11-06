"use client"
import { ArrowRightToLine, Send } from "lucide-react"
import { useState } from "react"
import Field from "../../ui/field/Field"

interface IMessageField {
	onSendMessage: (message: string) => void
}

export function MessageField({ onSendMessage }: IMessageField) {
	const [message, setMessage] = useState<string>("")

	const onSubmit = () => {
		if (message.trim()) {
			onSendMessage(message.trim())
			setMessage("")
		}
	}

	const handleKeyDown = (e: React.KeyboardEvent<HTMLTextAreaElement>) => {
		if (e.key === "Enter" && !e.shiftKey) {
			e.preventDefault()
			onSubmit()
		}
	}

	return (
		<div className='border-t border-border p-layout flex items-center w-full justify-between bottom-0'>
			<Field
				className='w-full'
				placeholder='Write a message...'
				Icon={ArrowRightToLine}
				value={message}
				onChange={(e) => setMessage(e.target.value)}
				onKeyDown={handleKeyDown}
				rows={1}
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
