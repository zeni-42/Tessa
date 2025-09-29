# Tessa

**A Clipboard Manager for Hyprland**

Tessa is a lightweight clipboard manager designed specifically for the Hyprland Wayland compositor. It tracks clipboard history and provides quick access to previously copied content, all while being minimal and efficient.

## Features

- Clipboard history tracking
- Designed for Hyprland (Wayland compositor)
- Minimal resource usage
- Simple and intuitive to use

## Installation

Make the install script executable and run it. The script will handle everything for you:

```bash
chmod +x scripts/install.sh
./scripts/install.sh
```

## Usage
Once running, Tessa monitors clipboard changes and maintains a history. Access your clipboard history through the application interface or your Hyprland configuration.

## Methord
- **save**

    Stores the provided data for later use. Everything after the save command is treated as a single string.
    ```bash
    tessa save [your data here]
    ```
- **show**

    Displays all previously saved data.
    ```bash
    tessa show
    ```
- **get**

    Retrieves a specific stored item by its index (as shown by the show command) and copies it for use.
    ```bash
    tessa get [index]
    ```
- **clean**

    Clears all previously stored data. Use with caution, as this cannot be undone.
    ```bash
    tessa clean
    ```
## Contributing
Contributions are welcome! Submit pull requests or open issues for bugs, enhancements, or suggestions.

# License
This project is licensed under the MIT License. See the [LICENSE](https://mit-license.org/)