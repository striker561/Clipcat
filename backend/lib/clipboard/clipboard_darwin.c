#include <Carbon/Carbon.h>

static EventHotKeyRef gHotKeyRef = NULL;

static OSStatus hotkeyHandler(EventHandlerCallRef nextHandler, EventRef event, void *userData) {
	extern void darwinHotkeyFired(void);
	darwinHotkeyFired();
	return noErr;
}

void registerDarwinHotkey(void) {
	EventTypeSpec eventType;
	eventType.eventClass = kEventClassKeyboard;
	eventType.eventKind = kEventHotKeyPressed;
	InstallApplicationEventHandler(NewEventHandlerUPP(hotkeyHandler), 1, &eventType, NULL, NULL);

	EventHotKeyID hotKeyID;
	hotKeyID.signature = 'CLIP';
	hotKeyID.id = 1;
	// kVK_ANSI_V = 9, Ctrl+Shift
	RegisterEventHotKey(9, controlKey | shiftKey, hotKeyID, GetApplicationEventTarget(), 0, &gHotKeyRef);
}

void unregisterDarwinHotkey(void) {
	if (gHotKeyRef != NULL) {
		UnregisterEventHotKey(gHotKeyRef);
		gHotKeyRef = NULL;
	}
}
