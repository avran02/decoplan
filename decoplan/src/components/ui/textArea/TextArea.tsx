import cn from "clsx"
import { forwardRef, useEffect, useRef } from "react"

import styles from "./textArea.module.scss"
import { TypeInputProps } from "./textArea.types"

const TextArea = forwardRef<HTMLTextAreaElement, TypeInputProps>(
	({ error, style, Icon, className, value, onChange, ...rest }, ref) => {
		const textareaRef = useRef<HTMLTextAreaElement>(null)

		useEffect(() => {
			const textarea = textareaRef.current
			if (textarea) {
				textarea.style.height = "42px"
				textarea.style.height = `${textarea.scrollHeight}px`
			}
		}, [value])

		return (
			<label className={cn(styles.field, className)} style={style}>
				{Icon && (
					<div className={styles.icon}>
						<Icon />
					</div>
				)}
				<textarea
					ref={(ref as any) || textareaRef}
					value={value}
					onChange={onChange}
					className={styles.textarea}
					{...rest}
				/>
				{error && <div className={styles.error}>{error.message}</div>}
			</label>
		)
	}
)

TextArea.displayName = "TextArea"

export default TextArea
