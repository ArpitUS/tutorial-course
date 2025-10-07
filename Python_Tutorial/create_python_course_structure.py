import os

# Base folder
base_dir = "python-course"

# Module names and titles
modules = [
    "Module_01_Basics",
    "Module_02_Variables",
    "Module_03_ControlFlow",
    "Module_04_Functions",
    "Module_05_DataStructures",
    "Module_06_StringsFiles",
    "Module_07_ErrorHandling",
    "Module_08_OOP",
    "Module_09_Modules",
    "Module_10_Automation",
]

# Top-level README.md
main_readme_content = """# 🐍 Python Tutorial Course — From Zero to Python Hero

Welcome to the **Python Tutorial Course**, designed for beginners who want to build strong foundations in Python programming through practical coding exercises.

This repository contains **10 structured modules**, each with:
- A short **README.md** explaining the topic
- A **main_before.py** file (starter code for practice)
- A **main_after.py** file (final reference solution)

---
Run each `main_before.py` to practice and compare with `main_after.py` for guidance.
"""

# Module README template
module_readme_template = """# {module_title}

### 🎯 Objectives
- Describe the topic briefly
- Write example code in `main_before.py`
- Compare results with `main_after.py`

---

### 🧠 Instructions
1. Open `main_before.py`
2. Follow the inline comments
3. Run the script
4. Review your solution with `main_after.py`

---

### ⏱ Duration
Approx. 5–6 minutes
"""

# Create base directory
os.makedirs(base_dir, exist_ok=True)

# Write main README.md
with open(os.path.join(base_dir, "README.md"), "w", encoding="utf-8") as f:
    f.write(main_readme_content)

# Create modules
for module in modules:
    module_path = os.path.join(base_dir, module)
    os.makedirs(module_path, exist_ok=True)

    # Create README.md for module
    readme_path = os.path.join(module_path, "README.md")
    with open(readme_path, "w", encoding="utf-8") as f:
        f.write(module_readme_template.format(module_title=module.replace("_", " ")))

    # Create empty main_before.py
    open(os.path.join(module_path, "main_before.py"), "w").write("# TODO: Write your code here\n")

    # Create empty main_after.py
    open(os.path.join(module_path, "main_after.py"), "w").write("# Reference solution\n")

print(f"✅ Python course structure created successfully at: {os.path.abspath(base_dir)}")
