"use client"
import { useEffect, useRef, useState } from "react"

const colors = ["#ff8800", "#e124ff", "#6a19ff", "#ff2188", "#f14848"]

export function NotFound() {
	const dvdRef = useRef<HTMLDivElement>(null)
	const [color, setColor] = useState<string>(colors[0])

	// Initialize animation variables
	let x = 0,
		y = 0,
		dirX = 1,
		dirY = 1
	const speed = 1
	const dvdWidth = 250
	const dvdHeight = 250
	let prevColorIndex = 0

	const getRandomColor = (): string => {
		const currentColors = [...colors]
		currentColors.splice(prevColorIndex, 1)
		const newIndex = Math.floor(Math.random() * currentColors.length)
		prevColorIndex = newIndex < prevColorIndex ? newIndex : newIndex + 1
		return currentColors[newIndex]
	}

	const animate = () => {
		const screenHeight = window.innerHeight
		const screenWidth = window.innerWidth

		if (y + dvdHeight >= screenHeight || y < 0) {
			dirY *= -1
			setColor(getRandomColor())
		}
		if (x + dvdWidth >= screenWidth || x < 0) {
			dirX *= -1
			setColor(getRandomColor())
		}

		// Update position
		x += dirX * speed
		y += dirY * speed

		// Apply transform style to DVD logo
		if (dvdRef.current) {
			dvdRef.current.style.transform = `translate(${x}px, ${y}px)`
		}

		requestAnimationFrame(animate)
	}

	useEffect(() => {
		requestAnimationFrame(animate)
	}, [])

	return (
		<div className='h-screen w-screen bg-black overflow-hidden'>
			<div
				ref={dvdRef}
				className='absolute'
				style={{
					width: dvdWidth,
					height: dvdHeight,
				}}
			>
				<svg
					xmlns='http://www.w3.org/2000/svg'
					className='icon icon-tabler icon-tabler-error-404'
					width={dvdWidth}
					height={dvdHeight}
					viewBox='0 0 24 24'
					strokeWidth='2'
					stroke={color}
					fill='none'
					strokeLinecap='round'
					strokeLinejoin='round'
				>
					<path stroke='none' d='M0 0h24v24H0z' fill='none'></path>
					<path d='M3 7v4a1 1 0 0 0 1 1h3'></path>
					<path d='M7 7v10'></path>
					<path d='M10 8v8a1 1 0 0 0 1 1h2a1 1 0 0 0 1 -1v-8a1 1 0 0 0 -1 -1h-2a1 1 0 0 0 -1 1z'></path>
					<path d='M17 7v4a1 1 0 0 0 1 1h3'></path>
					<path d='M21 7v10'></path>
				</svg>
			</div>
		</div>
	)
}
