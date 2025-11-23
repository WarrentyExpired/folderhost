import SwitchToggle from "../SwitchToggle/SwitchToggle"

interface PermissionToggleProps {
  label: string
  checked: boolean
  onChange: (checked: boolean) => void
}

const PermissionToggle = ({ label, checked, onChange }: PermissionToggleProps) => {
  return (
    <div className="flex items-center justify-between p-2 bg-gray-700 rounded transition-all group">
      <span className="text-white text-base group-hover:translate-x-1 transition-all">{label}</span>
      <SwitchToggle checked={checked} onChange={onChange} />
    </div>
  )
}

export default PermissionToggle;