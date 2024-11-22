# Logging libs

- [Logging libs](#logging-libs)
    - [1. Logrus](#1-logrus)
    - [2. Zap](#2-zap)
    - [3. Zerolog](#3-zerolog)
    - [Podsumowanie Wydajności](#podsumowanie-wydajności)
    - [Wybór Biblioteki](#wybór-biblioteki)


Wybór biblioteki do logowania, która jest mniej zasobożerna i bardziej wydajna, może zależeć od kilku czynników, takich jak konkretne wymagania projektu, sposób, w jaki logi są generowane, oraz jak są później przetwarzane. Poniżej przedstawiam porównanie trzech popularnych bibliotek logujących w Go: **Logrus**, **Zap**, i **Zerolog**.

### 1. Logrus

- **Wydajność**: Logrus jest prosty i elastyczny, ale jest też stosunkowo mniej wydajny w porównaniu do Zap i Zerolog. Jego narzut na logowanie może być większy, zwłaszcza przy dużych ilościach logów.
- **Zasobożerność**: W przypadku Logrus, jeśli logujesz dużo informacji w poziomie DEBUG, może to prowadzić do większego zużycia zasobów.
- **Formatowanie**: Domyślnie Logrus używa formatu tekstowego, ale może być skonfigurowany do używania JSON.

### 2. Zap

- **Wydajność**: Zap jest znany z wysokiej wydajności, dzięki efektywnemu zarządzaniu pamięcią i minimalnemu narzutowi. Jest zaprojektowany z myślą o wydajności i produkcyjnych zastosowaniach.
- **Zasobożerność**: Zap wykorzystuje lazy evaluation, co oznacza, że oblicza dane logów tylko wtedy, gdy są one rzeczywiście potrzebne. Oznacza to, że przy wyłączonym logowaniu na niższych poziomach (np. DEBUG), nie ma kosztów związanych z formatowaniem i konstruowaniem komunikatów logów.
- **Formatowanie**: Obsługuje format JSON i tekstowy, z domyślnym formatem JSON w trybie produkcyjnym.

### 3. Zerolog

- **Wydajność**: Zerolog jest jedną z najwydajniejszych bibliotek logujących dla Go, specjalizującą się w minimalnym narzucie i dużej szybkości. Jest zoptymalizowany pod kątem działania w systemach o wysokiej wydajności.
- **Zasobożerność**: Dzięki lazy evaluation oraz minimalnym operacjom podczas rejestrowania, Zerolog potrafi zminimalizować zużycie pamięci. Tworzy logi w formacie JSON w sposób bardzo wydajny.
- **Formatowanie**: Domyślnie logi są w formacie JSON, co jest korzystne dla systemów, które muszą przetwarzać logi w formacie strukturalnym.

### Podsumowanie Wydajności

- **Najlepsza wydajność**: **Zerolog** - zaprojektowany z myślą o dużej wydajności i niskim narzucie zasobów. Idealny do środowisk produkcyjnych, gdzie ilość logów jest wysoka.
- **Wysoka wydajność**: **Zap** - bardzo wydajny, ale może mieć nieco wyższy narzut niż Zerolog w niektórych scenariuszach.
- **Mniejsza wydajność**: **Logrus** - oferuje prostotę i elastyczność, ale może być mniej wydajny w aplikacjach o dużym obciążeniu.

### Wybór Biblioteki

Jeśli Twoim głównym celem jest wydajność i minimalne zużycie zasobów, **Zerolog** jest najlepszym wyborem. Jeśli jednak potrzebujesz bardziej zaawansowanych funkcji i nieco większej elastyczności, **Zap** również będzie bardzo dobrym rozwiązaniem. **Logrus** jest świetny dla prostszych zastosowań, gdzie wydajność nie jest kluczowym czynnikiem. 

Ostateczny wybór powinien być dostosowany do potrzeb Twojej aplikacji, obciążenia, jakiego się spodziewasz, oraz sposobu, w jaki planujesz przetwarzać logi.
