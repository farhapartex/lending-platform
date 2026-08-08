export const onboardingCompletedKey = "lendingPlatform.onboardingCompleted";

export function markOnboardingComplete(): void {
  try {
    window.localStorage.setItem(onboardingCompletedKey, "true");
  } catch {
    return;
  }
}
