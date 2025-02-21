package w32uiautomation

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IUIAutomationTreeWalker struct {
	ole.IUnknown
}

type IUIAutomationTreeWalkerVtbl struct {
	ole.IUnknownVtbl
	GetParentElement                    uintptr
	GetFirstChildElement                uintptr
	GetLastChildElement                 uintptr
	GetNextSiblingElement               uintptr
	GetPreviousSiblingElement           uintptr
	NormalizeElement                    uintptr
	GetParentElementBuildCache          uintptr
	GetFirstChildElementBuildCache      uintptr
	GetLastChildElementBuildCache       uintptr
	GetNextSiblingElementBuildCache     uintptr
	GetPreviousSiblingElementBuildCache uintptr
	NormalizeElementBuildCache          uintptr
	Get_Condition                       uintptr
}

func (w *IUIAutomationTreeWalker) VTable() *IUIAutomationTreeWalkerVtbl {
	return (*IUIAutomationTreeWalkerVtbl)(unsafe.Pointer(w.RawVTable))
}

func (w *IUIAutomationTreeWalker) GetParentElement(element *IUIAutomationElement) (first *IUIAutomationElement, err error) {
	return getParentElement(w, element)
}

func (w *IUIAutomationTreeWalker) GetFirstChildElement(element *IUIAutomationElement) (first *IUIAutomationElement, err error) {
	return getFirstChildElement(w, element)
}

func (w *IUIAutomationTreeWalker) GetNextSiblingElement(element *IUIAutomationElement) (next *IUIAutomationElement, err error) {
	return getNextSiblingElement(w, element)
}

func (w *IUIAutomationTreeWalker) GetPreviousSiblingElement(element *IUIAutomationElement) (previous *IUIAutomationElement, err error) {
	return getPreviousSiblingElement(w, element)
}

func getParentElement(w *IUIAutomationTreeWalker, element *IUIAutomationElement) (parent *IUIAutomationElement, err error) {
	hr, _, _ := syscall.SyscallN(
		w.VTable().GetParentElement,
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(element)),
		uintptr(unsafe.Pointer(&parent)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func getFirstChildElement(w *IUIAutomationTreeWalker, element *IUIAutomationElement) (first *IUIAutomationElement, err error) {
	hr, _, _ := syscall.SyscallN(
		w.VTable().GetFirstChildElement,
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(element)),
		uintptr(unsafe.Pointer(&first)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func getNextSiblingElement(w *IUIAutomationTreeWalker, element *IUIAutomationElement) (next *IUIAutomationElement, err error) {
	hr, _, _ := syscall.SyscallN(
		w.VTable().GetNextSiblingElement,
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(element)),
		uintptr(unsafe.Pointer(&next)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func getPreviousSiblingElement(w *IUIAutomationTreeWalker, element *IUIAutomationElement) (previous *IUIAutomationElement, err error) {
	hr, _, _ := syscall.SyscallN(
		w.VTable().GetPreviousSiblingElement,
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(element)),
		uintptr(unsafe.Pointer(&previous)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
