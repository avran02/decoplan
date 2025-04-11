"use client"

import { IChat } from "@/types/chat.types"
import Image from "next/image"
import Link from "next/link"

interface IChatListItem {
	chat: IChat
}

export function ChatListItem({ chat }: IChatListItem) {
	return (
		<Link
			href={`/chat/${chat.id}`}
			className='p-layout flex items-center border-b border-border duration-300 ease-linear transition-colors hover:bg-border animation-slide-fade'
		>
			<Image
				src={chat.avatar || "/no-avatar.png"}
				alt={""}
				width={45}
				height={45}
				className='mr-4'
			/>
			<div className='text-sm w-full'>
				<div className='flex items-center justify-between'>
					<span>{chat.chatName}</span>
				</div>
			</div>
		</Link>
	)
}
