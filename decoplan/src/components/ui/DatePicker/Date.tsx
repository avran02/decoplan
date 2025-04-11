import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs"
import { DatePicker } from "@mui/x-date-pickers/DatePicker"
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider"
import { Controller } from "react-hook-form"

interface DateProps {
	control: any
	defaultValue?: Date | null
}

export function Date({ control, defaultValue }: DateProps) {
	return (
		<LocalizationProvider dateAdapter={AdapterDayjs}>
			<Controller
				name='birthDate'
				control={control}
				rules={{ required: true }}
				defaultValue={defaultValue}
				render={({ field }) => (
					<DatePicker
						{...field}
						value={field.value || null}
						sx={{
							width: "100%",
							"& .MuiInputLabel-root": {
								color: "rgb(75, 85, 99, 0.9)",
								borderColor: "rgb(64 64 64 / 1)",
							},
							"& .MuiOutlinedInput-input": {
								color: "gray",
								border: "none",
							},
							"&:hover &.MuiOutlinedInput-input": {
								border: "none",
							},
							"& .MuiInputBase-adornedEnd": {
								border: "1.5px solid rgba(255, 255, 255, 0.12)",
								borderRadius: "0.25rem",
							},
							"&:hover .MuiInputBase-adornedEnd": {
								border: "1.5px solid rgba(255, 255, 255, 0.12)",
								borderRadius: "0.25rem",
							},
							"&:focus .Mui-focused": {
								border: "rgba(99, 102, 241, 1)",
								borderRadius: "0.25rem",
							},
							"& .MuiIconButton-root": {
								color: "rgb(75, 85, 99, 0.9)",
							},
						}}
						format='DD/MM/YYYY'
						slotProps={{ field: { shouldRespectLeadingZeros: true } }}
					/>
				)}
			/>
		</LocalizationProvider>
	)
}
