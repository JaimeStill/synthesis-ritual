# Cosmodrome Design System

Cosmodrome is a comprehensive color system designed for applications with a retro-futuristic and space exploration aesthetic. Drawing inspiration from vintage NASA control rooms, technical blueprints, and retro sci-fi interfaces, Cosmodrome provides a balanced approach to color with optimized accessibility and visual harmony. Initially conceived as a spin of [Catppuccin](https://catppuccin.com/).

## Theme Flavors

Cosmodrome offers two distinct palette flavors, each capturing a different facet of space exploration aesthetics:

1. **Apollo**: A dark theme combining the warm, vintage feel of Apollo-era mission control rooms with vibrant accent colors inspired by retro sci-fi interfaces and NASA control panels

2. **Horizon**: A light theme merging the technical precision of engineering blueprints with the optimistic palette of Space Age promotional materials and retrofuturistic illustrations

## Apollo (Dark Theme)

A dark theme combining the warm, vintage feel of Apollo-era mission control rooms with vibrant accent colors inspired by retro sci-fi interfaces.

### Base Colors

| Color | Value | Description |
|-------|-------|-------------|
| Background | `#1A1D21` | Dark charcoal with warm undertone |
| Surface | `#25292E` | Slightly elevated surface color |
| Overlay | `#303540` | Surface for overlays and popups |
| Text | `#E8E3DC` | Warm off-white primary text |
| Subtext | `#B6AEA1` | Muted warm gray secondary text |

### Accent & Contrast Colors

| Color | Accent Value | Contrast Value | Description |
|-------|-------------|---------------|-------------|
| Red | `#FF5A5F` | `#D13034` | Retro alarm red |
| Orange | `#FF9E64` | `#D57339` | Vintage equipment orange |
| Yellow | `#FFCB6B` | `#D9A040` | Warning console yellow |
| Green | `#7BD88F` | `#50AD64` | Terminal green |
| Teal | `#5CCCC6` | `#31A19B` | Vintage monitor teal |
| Blue | `#69A7EB` | `#3E7CC0` | Control panel blue |
| Purple | `#BF94E4` | `#9469B9` | Sci-fi purple |
| Pink | `#F0A6CA` | `#C57B9F` | Retro futurism pink |

### Accessibility Notes

- Text on Background: 13.25:1 ratio (Excellent)
- Subtext on Background: 7.70:1 ratio (Excellent)
- Most accent colors meet AA standards against dark backgrounds
- Contrast colors ensure accessibility for interactive elements

### Color Variables (CSS)

```css
/* Apollo Theme */
:root {
  /* Base */
  --apollo-bg: #1A1D21;
  --apollo-surface: #25292E;
  --apollo-overlay: #303540;
  --apollo-text: #E8E3DC;
  --apollo-subtext: #B6AEA1;
  
  /* Accent */
  --apollo-red: #FF5A5F;
  --apollo-orange: #FF9E64;
  --apollo-yellow: #FFCB6B;
  --apollo-green: #7BD88F;
  --apollo-teal: #5CCCC6;
  --apollo-blue: #69A7EB;
  --apollo-purple: #BF94E4;
  --apollo-pink: #F0A6CA;
  
  /* Contrast */
  --apollo-red-contrast: #D13034;
  --apollo-orange-contrast: #D57339;
  --apollo-yellow-contrast: #D9A040;
  --apollo-green-contrast: #50AD64;
  --apollo-teal-contrast: #31A19B;
  --apollo-blue-contrast: #3E7CC0;
  --apollo-purple-contrast: #9469B9;
  --apollo-pink-contrast: #C57B9F;
}
```

## Horizon (Light Theme)

A light theme merging the technical precision of engineering blueprints with the optimistic palette of Space Age promotional materials.

### Base Colors

| Color | Value | Description |
|-------|-------|-------------|
| Background | `#D5D8DB` | Light gray with slight blue tint |
| Surface | `#C5CAD3` | Slightly darker surface with blueprint feel |
| Overlay | `#B7BEC9` | Overlay and elevated UI elements |
| Text | `#242D3A` | Deep blue-black primary text |
| Subtext | `#536171` | Medium blue-gray secondary text |

### Accent & Contrast Colors

| Color | Accent Value | Contrast Value | Description |
|-------|-------------|---------------|-------------|
| Red | `#CF3B43` | `#901923` | Atomic Age red |
| Orange | `#D7623B` | `#983114` | Retrofuture orange |
| Yellow | `#E19642` | `#A26421` | Space Age yellow |
| Green | `#398060` | `#1F4B36` | Futuristic vision green |
| Teal | `#2F8985` | `#154B48` | Concept art teal |
| Blue | `#2C5F9B` | `#11305A` | Utopian future blue |
| Purple | `#6F5697` | `#3F2B5F` | Sci-fi magazine purple |
| Pink | `#B34A6D` | `#742337` | Space colony pink |

### Accessibility Notes

- Text on Background: 9.71:1 ratio (Excellent)
- Subtext on Background: 4.43:1 ratio (Good, meets AA for large text)
- Accent colors have good contrast with light background
- Contrast colors provide excellent contrast for interactive states

### Color Variables (CSS)

```css
/* Horizon Theme */
:root {
  /* Base */
  --horizon-bg: #D5D8DB;
  --horizon-surface: #C5CAD3;
  --horizon-overlay: #B7BEC9;
  --horizon-text: #242D3A;
  --horizon-subtext: #536171;
  
  /* Accent */
  --horizon-red: #CF3B43;
  --horizon-orange: #D7623B;
  --horizon-yellow: #E19642;
  --horizon-green: #398060;
  --horizon-teal: #2F8985;
  --horizon-blue: #2C5F9B;
  --horizon-purple: #6F5697;
  --horizon-pink: #B34A6D;
  
  /* Contrast */
  --horizon-red-contrast: #901923;
  --horizon-orange-contrast: #983114;
  --horizon-yellow-contrast: #A26421;
  --horizon-green-contrast: #1F4B36;
  --horizon-teal-contrast: #154B48;
  --horizon-blue-contrast: #11305A;
  --horizon-purple-contrast: #3F2B5F;
  --horizon-pink-contrast: #742337;
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

## Accessibility

All color combinations meet or exceed WCAG 2.1 AA standards for text legibility:

- Normal text (smaller than 18pt): Minimum contrast ratio of 4.5:1
- Large text (larger than 18pt or bold): Minimum contrast ratio of 3:1

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
