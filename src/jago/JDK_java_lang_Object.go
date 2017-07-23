package jago

func register_java_lang_Object() {
	register("java/lang/Object.registerNatives()V", Java_jang_lang_Object_registerNatives)
	register("java/lang/Object.hashCode()I", Java_java_lang_Object_hashCode)
	register("java/lang/Object.getClass()Ljava/lang/Class;", Java_java_lang_Object_getClass)
	register("java/lang/Object.clone()Ljava/lang/Object;", Java_java_lang_Object_clone)
}

// private static void registerNatives()
func Java_jang_lang_Object_registerNatives()  {}

func Java_java_lang_Object_hashCode(this Reference) Int  {
	return VM_iHashCode(this)
}

func Java_java_lang_Object_getClass(this Reference) JavaLangClass {
	return this.Class().ClassObject()
}

func Java_java_lang_Object_clone(this Reference) Reference {
	cloneable := BOOTSTRAP_CLASSLOADER.CreateClass("java/lang/Cloneable")
	if !cloneable.IsAssignableFrom(this.Class()) {
		panic("java.lang.CloneNotSupportedException")
	}

	return __clone(this)
}

func __clone(obj Reference) Reference {

	var clone Reference
	if !obj.Class().IsArray() {
		clone = obj.Class().NewObject().(Reference)
	} else {
		clone = obj.Class().NewArray(obj.Length()).(Reference)
	}

	for i,value := range obj.oop.values {
		clone.oop.values[i] = value
	}

	return clone
}
