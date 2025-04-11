"use client"

import cn from "clsx"
import { Sun } from "lucide-react"
import Link from "next/link"
import { usePathname } from "next/navigation"

import { MENU } from "./sidebar.data"

import { useAuth } from "@/hooks/useAuth"
import { useEffect, useState } from "react"
import styles from "./Sidebar.module.scss"

export function Sidebar() {
	const { isLoggedIn } = useAuth()
	const pathname = usePathname()

	const [isClient, setIsClient] = useState(false)

	useEffect(() => {
		setIsClient(true)
	}, [])

	if (!isClient) return null

	return (
		<aside className={cn(styles.sidebar, { [styles.loggedIn]: isLoggedIn })}>
			{isLoggedIn ? (
				<>
					{/* <Image src='/logo.svg' priority alt='' width={45} height={45} /> */}
					<div>
						{MENU.map((item) => (
							<Link
								href={item.url}
								key={item.url}
								className={cn({
									[styles.active]: pathname === item.url,
								})}
							>
								<item.icon size={27} />
							</Link>
						))}
					</div>
					<Sun size={27} />
				</>
			) : null}
		</aside>
	)
}
