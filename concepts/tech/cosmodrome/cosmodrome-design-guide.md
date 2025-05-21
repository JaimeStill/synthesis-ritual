# Cosmodrome Design System

Cosmodrome is a comprehensive color system designed for applications with a retro-futuristic and space exploration aesthetic. Drawing inspiration from vintage NASA control rooms, technical blueprints, and retro sci-fi interfaces, Cosmodrome provides a balanced approach to color with optimized accessibility and visual harmony. Conceived as a spin of [Catppuccin](https://catppuccin.com/).

## Theme Flavors

Cosmodrome offers two distinct palette flavors, each capturing a different facet of space exploration aesthetics:

1. **Apollo**: A dark theme combining the warm, vintage feel of Apollo-era mission control rooms with vibrant accent colors inspired by retro sci-fi interfaces and NASA control panels

2. **Horizon**: A light theme merging the technical precision of engineering blueprints with the optimistic palette of Space Age promotional materials and retrofuturistic illustrations

## Palette Structure

Each flavor includes 13 standardized colors:

- **Infrastructure** (3): Base, Surface, Overlay
- **Text** (2): Text, Subtext
- **Primary Colors** (8): Red, Orange, Yellow, Green, Teal, Blue, Purple, Pink

Every color in the sub-palette has a main version and a contrast version (except Text/Subtext).

## Infrastructure

Infrastructure refers to the hierarchical background layers of the UI:

- **Base**: The primary background color that serves as the foundation
- **Surface**: Slightly elevated elements that sit on top of the base layer
- **Overlay**: Higher elevation elements like modals, popups, and dialogs

## Apollo (Dark Theme)

A dark theme combining the warm, vintage feel of Apollo-era mission control rooms with vibrant accent colors inspired by retro sci-fi interfaces.

### Base Colors

| Color | Value | RGB | HSL | Description |
|-------|-------|-----|-----|-------------|
| Base | `#1A1D21` | `26, 29, 33` | `214, 12%, 12%` | Dark charcoal with warm undertone |
| Surface | `#25292E` | `37, 41, 46` | `214, 11%, 16%` | Slightly elevated surface color |
| Overlay | `#303540` | `48, 53, 64` | `220, 14%, 22%` | Surface for overlays and popups |
| Text | `#E8E3DC` | `232, 227, 220` | `39, 24%, 89%` | Warm off-white primary text |
| Subtext | `#B6AEA1` | `182, 174, 161` | `39, 13%, 67%` | Muted warm gray secondary text |

### Accent & Contrast Colors

| Color | Accent Value | Contrast Value | RGB (Accent) | RGB (Contrast) | Description |
|-------|-------------|---------------|-------------|---------------|-------------|
| Red | `#FF5A5F` | `#D13034` | `255, 90, 95` | `209, 48, 52` | Retro alarm red |
| Orange | `#FF9E64` | `#D57339` | `255, 158, 100` | `213, 115, 57` | Vintage equipment orange |
| Yellow | `#FFCB6B` | `#D9A040` | `255, 203, 107` | `217, 160, 64` | Warning console yellow |
| Green | `#7BD88F` | `#50AD64` | `123, 216, 143` | `80, 173, 100` | Terminal green |
| Teal | `#5CCCC6` | `#31A19B` | `92, 204, 198` | `49, 161, 155` | Vintage monitor teal |
| Blue | `#69A7EB` | `#3E7CC0` | `105, 167, 235` | `62, 124, 192` | Control panel blue |
| Purple | `#BF94E4` | `#9469B9` | `191, 148, 228` | `148, 105, 185` | Sci-fi purple |
| Pink | `#F0A6CA` | `#C57B9F` | `240, 166, 202` | `197, 123, 159` | Retro futurism pink |

### Accessibility Notes

- Text on Base: 13.25:1 ratio (Excellent)
- Subtext on Base: 7.70:1 ratio (Excellent)
- Most accent colors meet AA standards against dark backgrounds
- Contrast colors ensure accessibility for interactive elements

### Color Variables (CSS)

```css
/* Apollo Theme */
:root {
  /* Base */
  --cosmodrome-base: #1A1D21;
  --cosmodrome-surface: #25292E;
  --cosmodrome-overlay: #303540;
  --cosmodrome-text: #E8E3DC;
  --cosmodrome-subtext: #B6AEA1;
  
  /* Accent */
  --cosmodrome-red: #FF5A5F;
  --cosmodrome-orange: #FF9E64;
  --cosmodrome-yellow: #FFCB6B;
  --cosmodrome-green: #7BD88F;
  --cosmodrome-teal: #5CCCC6;
  --cosmodrome-blue: #69A7EB;
  --cosmodrome-purple: #BF94E4;
  --cosmodrome-pink: #F0A6CA;
  
  /* Contrast */
  --cosmodrome-red-contrast: #D13034;
  --cosmodrome-orange-contrast: #D57339;
  --cosmodrome-yellow-contrast: #D9A040;
  --cosmodrome-green-contrast: #50AD64;
  --cosmodrome-teal-contrast: #31A19B;
  --cosmodrome-blue-contrast: #3E7CC0;
  --cosmodrome-purple-contrast: #9469B9;
  --cosmodrome-pink-contrast: #C57B9F;
}
```

## Horizon (Light Theme)

A light theme merging the technical precision of engineering blueprints with the optimistic palette of Space Age promotional materials.

### Base Colors

| Color | Value | RGB | HSL | Description |
|-------|-------|-----|-----|-------------|
| Base | `#D5D8DB` | `213, 216, 219` | `210, 7%, 85%` | Light gray with slight blue tint |
| Surface | `#C5CAD3` | `197, 202, 211` | `217, 14%, 80%` | Slightly darker surface with blueprint feel |
| Overlay | `#B7BEC9` | `183, 190, 201` | `217, 15%, 75%` | Overlay and elevated UI elements |
| Text | `#242D3A` | `36, 45, 58` | `217, 23%, 18%` | Deep blue-black primary text |
| Subtext | `#536171` | `83, 97, 113` | `217, 15%, 38%` | Medium blue-gray secondary text |

### Accent & Contrast Colors

| Color | Accent Value | Contrast Value | RGB (Accent) | RGB (Contrast) | Description |
|-------|-------------|---------------|-------------|---------------|-------------|
| Red | `#CF3B43` | `#901923` | `207, 59, 67` | `144, 25, 35` | Atomic Age red |
| Orange | `#D7623B` | `#983114` | `215, 98, 59` | `152, 49, 20` | Retrofuture orange |
| Yellow | `#E19642` | `#A26421` | `225, 150, 66` | `162, 100, 33` | Space Age yellow |
| Green | `#398060` | `#1F4B36` | `57, 128, 96` | `31, 75, 54` | Futuristic vision green |
| Teal | `#2F8985` | `#154B48` | `47, 137, 133` | `21, 75, 72` | Concept art teal |
| Blue | `#2C5F9B` | `#11305A` | `44, 95, 155` | `17, 48, 90` | Utopian future blue |
| Purple | `#6F5697` | `#3F2B5F` | `111, 86, 151` | `63, 43, 95` | Sci-fi magazine purple |
| Pink | `#B34A6D` | `#742337` | `179, 74, 109` | `116, 35, 55` | Space colony pink |

### Accessibility Notes

- Text on Base: 9.71:1 ratio (Excellent)
- Subtext on Base: 4.43:1 ratio (Good, meets AA for large text)
- Accent colors have good contrast with light background
- Contrast colors provide excellent contrast for interactive states

### Color Variables (CSS)

```css
/* Horizon Theme */
:root {
  /* Base */
  --cosmodrome-base: #D5D8DB;
  --cosmodrome-surface: #C5CAD3;
  --cosmodrome-overlay: #B7BEC9;
  --cosmodrome-text: #242D3A;
  --cosmodrome-subtext: #536171;
  
  /* Accent */
  --cosmodrome-red: #CF3B43;
  --cosmodrome-orange: #D7623B;
  --cosmodrome-yellow: #E19642;
  --cosmodrome-green: #398060;
  --cosmodrome-teal: #2F8985;
  --cosmodrome-blue: #2C5F9B;
  --cosmodrome-purple: #6F5697;
  --cosmodrome-pink: #B34A6D;
  
  /* Contrast */
  --cosmodrome-red-contrast: #901923;
  --cosmodrome-orange-contrast: #983114;
  --cosmodrome-yellow-contrast: #A26421;
  --cosmodrome-green-contrast: #1F4B36;
  --cosmodrome-teal-contrast: #154B48;
  --cosmodrome-blue-contrast: #11305A;
  --cosmodrome-purple-contrast: #3F2B5F;
  --cosmodrome-pink-contrast: #742337;
}
```

## Usage & Implementation Guidelines

Each palette includes specific color types for consistent implementation:

- **Base Colors**: Use for backgrounds, surfaces, content areas, and text
- **Accent Colors**: Use for interactive elements, highlights, and decorative components
- **Contrast Colors**: Darker variants of accent colors for states (hover, pressed) and to ensure accessibility

When implementing Cosmodrome in your applications:

1. **Select a flavor** that best matches your application's mood and target environment
2. **Use base colors** for structural elements (backgrounds, containers, text)
3. **Apply accent colors** for interactive elements, highlights, and decorative components
4. **Use contrast colors** for:
   - Hover/focus/active states
   - Borders and outlines
   - Text on accent-colored backgrounds
   - Ensuring accessible contrast ratios

### Theme Implementation

The recommended approach for implementing themes:

1. **Separate CSS files** - Each theme should be encapsulated in its own CSS file
2. **Conditional loading** - Use media queries to load the appropriate theme:

```css
@import './apollo.css' screen;
@import './horizon.css' screen and (prefers-color-scheme: light);
```

This approach allows the system to automatically switch between themes based on the user's system preferences while making it easy to override with custom selectors if needed.

## Style Guide

### Typography

| Function | Element | Apollo | Horizon |
|----------|---------|--------|---------|
| Headings | h1, h2, h3 | Text on Base | Text on Base |
| Body | p, li, td | Text on Base | Text on Base |
| Secondary Text | captions, help text | Subtext on Base | Subtext on Base |
| Code | pre, code | Text on Surface | Text on Surface |
| Links, URLs | a | Blue on Base | Blue on Base |
| Link (Hover) | a:hover | Blue (Contrast) on Base | Blue (Contrast) on Base |
| Link (Visited) | a:visited | Purple on Base | Purple on Base |

### UI Elements

#### Buttons & Interactive Controls

| Function | Element | Apollo | Horizon |
|----------|---------|--------|---------|
| Primary Action | buttons, submit | Teal with Text | Teal with Base |
| Primary (Hover) | buttons:hover | Teal (Contrast) with Text | Teal (Contrast) with Base |
| Secondary Action | buttons\[secondary] | Surface with Text | Surface with Text |
| Secondary (Hover) | buttons\[secondary]:hover | Overlay with Text | Overlay with Text |
| Destructive Action | buttons\[destructive] | Red with Text | Red with Base |
| Destructive (Hover) | buttons\[destructive]:hover | Red (Contrast) with Text | Red (Contrast) with Base |
| Disabled | buttons\[disabled] | Surface with Subtext (opacity 0.5) | Surface with Subtext (opacity 0.5) |

#### Form Elements

| Function | Element | Apollo | Horizon |
|----------|---------|--------|---------|
| Input Background | input, select, textarea | Surface | Surface |
| Input Border | input, select, textarea | Overlay | Overlay |
| Input Border (Focus) | input:focus | Blue | Blue |
| Input Text | input, select, textarea | Text | Text |
| Placeholder | ::placeholder | Subtext (opacity 0.7) | Subtext (opacity 0.7) |
| Checkbox, Radio (Unchecked) | input\[type=checkbox] | Surface with Overlay border | Surface with Overlay border |
| Checkbox, Radio (Checked) | input\[type=checkbox]:checked | Green with Base | Green with Base |

#### Notifications & Alerts

| Function | Element | Apollo | Horizon |
|----------|---------|--------|---------|
| Info | .info, .notice | Blue (opacity 0.15) with Blue text | Blue (opacity 0.15) with Blue (Contrast) text |
| Success | .success | Green (opacity 0.15) with Green text | Green (opacity 0.15) with Green (Contrast) text |
| Warning | .warning | Yellow (opacity 0.15) with Yellow (Contrast) text | Yellow (opacity 0.15) with Yellow (Contrast) text |
| Error | .error, .danger | Red (opacity 0.15) with Red text | Red (opacity 0.15) with Red (Contrast) text |
| Subtle Alert | .alert-subtle | Surface with Subtext | Surface with Subtext |

#### Navigation & Menus

| Function | Element | Apollo | Horizon |
|----------|---------|--------|---------|
| Nav Background | nav, header | Surface | Surface |
| Nav Item | nav a | Text | Text |
| Nav Item (Active) | nav a.active | Teal | Teal (Contrast) |
| Nav Item (Hover) | nav a:hover | Text with Teal underline | Text with Teal (Contrast) underline |
| Dropdown Menu | .dropdown-menu | Overlay | Overlay |
| Dropdown Item | .dropdown-item | Text | Text |
| Dropdown Item (Hover) | .dropdown-item:hover | Surface with Text | Surface with Text |
| Divider | hr, .divider | Overlay (opacity 0.5) | Overlay (opacity 0.5) |

#### Status Indicators

| Function | Element | Apollo | Horizon |
|----------|---------|--------|---------|
| Online/Available | .status-online | Green | Green |
| Away/Inactive | .status-away | Yellow | Yellow |
| Busy/DND | .status-busy | Orange | Orange |
| Offline/Unavailable | .status-offline | Subtext | Subtext |
| New/Unread | .status-new | Blue | Blue |
| Loading | .loading, .spinner | Purple | Purple |

### Syntax Highlighting

| Function | Element | Apollo | Horizon |
|----------|---------|--------|---------|
| Background | pre, code | Surface | Surface |
| Plain Text | text | Text | Text |
| Comments | comment | Subtext | Subtext |
| Keywords | keyword | Purple | Purple (Contrast) |
| Strings | string | Green | Green (Contrast) |
| Numbers | number | Orange | Orange (Contrast) |
| Functions | function | Blue | Blue (Contrast) |
| Classes/Types | class, type | Yellow | Yellow (Contrast) |
| Attributes | attribute | Pink | Pink (Contrast) |
| Variables | variable | Teal | Teal (Contrast) |
| Tags | tag | Red | Red (Contrast) |

## Semantic Usage

Beyond theming, Cosmodrome colors can be used semantically:

| Purpose | Apollo (Dark) | Horizon (Light) |
|---------|---------------|-----------------|
| Success | Green | Green |
| Warning | Yellow/Orange | Yellow/Orange |
| Error | Red | Red |
| Info | Blue | Blue |
| Selected | Purple | Purple |
| New/Notable | Pink/Teal | Pink/Teal |

## Accessibility

All color combinations meet or exceed WCAG 2.1 AA standards for text legibility:

- Normal text (smaller than 18pt): Minimum contrast ratio of 4.5:1
- Large text (larger than 18pt or bold): Minimum contrast ratio of 3:1

## Font Recommendations

While Cosmodrome focuses on color, these font pairings complement the aesthetic:

- **Display/Headings**: Eurostile, Orbitron, Microgramma, Chakra Petch
- **Body/Code**: Space Mono, JetBrains Mono, IBM Plex Mono, Source Code Pro
- **UI Elements**: Inter, Public Sans, Work Sans

## Design Elements

Consider these design elements to enhance the retro futurism aesthetic:

1. **Grid patterns**: Blueprint-style grids or terminal-like scanlines
2. **Subtle gradients**: Especially in the dark theme
3. **Terminal-style animations**: For state changes and transitions
4. **Space-inspired iconography**: Stars, orbits, retrofuturistic symbols
5. **NASA-inspired interface elements**: Mission patches, technical diagrams
