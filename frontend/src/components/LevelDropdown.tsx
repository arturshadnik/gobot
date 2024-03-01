import React from "react";
import { Select, MenuItem } from "@mui/material";

type DropdownProps = {
    levels_list: string[];
    level: string;
    handleLevelChange: (level: string) => void;
}

const LevelDropdown: React.FC<DropdownProps> = ({ levels_list, level, handleLevelChange }) => {
    return (
        <Select
            value={level}
            label="level"
            onChange={(e) => handleLevelChange(e.target.value)}
            sx={{ color: "inherit", p: 0, borderRadius: "1rem", height: "45px" }}
        >
            {levels_list.map((levelElement: string) => (
                <MenuItem
                    key={levelElement}
                    value={levelElement}
                    sx={{ display: "flex", alignItems: "center" }}
                >
                    {levelElement}
                </MenuItem>
            ))}
        </Select>
    )
}

export default LevelDropdown;