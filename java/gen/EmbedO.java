package gen;


// Code generated by colf(1); DO NOT EDIT.
// The compiler used schema file test.colf.


import static java.lang.String.format;
import java.io.IOException;
import java.io.InputStream;
import java.io.ObjectInputStream;
import java.io.ObjectOutputStream;
import java.io.ObjectStreamException;
import java.io.OutputStream;
import java.io.Serializable;
import java.util.InputMismatchException;
import java.nio.BufferOverflowException;
import java.nio.BufferUnderflowException;


/**
 * Data bean with built-in serialization support.
 * EmbedO has an inner object only.
 * Covers regression of issue #66.
 * @author generated by colf(1)
 * @see <a href="https://github.com/frbrno/colfer">Colfer's home</a>
 */
public class EmbedO implements Serializable {

	/** The upper limit for serial byte sizes. */
	public static int colferSizeMax = 16 * 1024 * 1024;


	public O inner;

	/** Default constructor */
	public EmbedO() {
		init();
	}


	/** Colfer zero values. */
	private void init() {
	}

	/**
	 * {@link #reset(InputStream) Reusable} deserialization of Colfer streams.
	 */
	public static class Unmarshaller {

		/** The data source. */
		protected InputStream in;

		/** The read buffer. */
		public byte[] buf;

		/** The {@link #buf buffer}'s data start index, inclusive. */
		protected int offset;

		/** The {@link #buf buffer}'s data end index, exclusive. */
		protected int i;


		/**
		 * @param in the data source or {@code null}.
		 * @param buf the initial buffer or {@code null}.
		 */
		public Unmarshaller(InputStream in, byte[] buf) {
			if (buf == null || buf.length == 0)
				buf = new byte[Math.min(EmbedO.colferSizeMax, 2048)];
			this.buf = buf;
			reset(in);
		}

		/**
		 * Reuses the marshaller.
		 * @param in the data source or {@code null}.
		 * @throws IllegalStateException on pending data.
		 */
		public void reset(InputStream in) {
			if (this.i != this.offset) throw new IllegalStateException("colfer: pending data");
			this.in = in;
			this.offset = 0;
			this.i = 0;
		}

		/**
		 * Deserializes the following object.
		 * @return the result or {@code null} when EOF.
		 * @throws IOException from the input stream.
		 * @throws SecurityException on an upper limit breach defined by {@link #colferSizeMax}.
		 * @throws InputMismatchException when the data does not match this object's schema.
		 */
		public EmbedO next() throws IOException {
			if (in == null) return null;

			while (true) {
				if (this.i > this.offset) {
					try {
						EmbedO o = new EmbedO();
						this.offset = o.unmarshal(this.buf, this.offset, this.i);
						return o;
					} catch (BufferUnderflowException e) {
					}
				}
				// not enough data

				if (this.i <= this.offset) {
					this.offset = 0;
					this.i = 0;
				} else if (i == buf.length) {
					byte[] src = this.buf;
					if (offset == 0) this.buf = new byte[Math.min(EmbedO.colferSizeMax, this.buf.length * 4)];
					System.arraycopy(src, this.offset, this.buf, 0, this.i - this.offset);
					this.i -= this.offset;
					this.offset = 0;
				}
				assert this.i < this.buf.length;

				int n = in.read(buf, i, buf.length - i);
				if (n < 0) {
					if (this.i > this.offset)
						throw new InputMismatchException("colfer: pending data with EOF");
					return null;
				}
				assert n > 0;
				i += n;
			}
		}

	}

	/**
	 * Gets the serial size estimate as an upper boundary, whereby
	 * {@link #marshal(byte[],int)} ≤ {@link #marshalFit()} ≤ {@link #colferSizeMax}.
	 * @return the number of bytes.
	 */
	public int marshalFit() {
		long n = 1L;
		if (this.inner != null) n += 1 + (long)this.inner.marshalFit();
		if (n < 0 || n > (long)EmbedO.colferSizeMax) return EmbedO.colferSizeMax;
		return (int) n;
	}

	/**
	 * Serializes the object.
	 * @param out the data destination.
	 * @param buf the initial buffer or {@code null}.
	 * @return the final buffer. When the serial fits into {@code buf} then the return is {@code buf}.
	 *  Otherwise the return is a new buffer, large enough to hold the whole serial.
	 * @throws IOException from {@code out}.
	 * @throws IllegalStateException on an upper limit breach defined by {@link #colferSizeMax}.
	 */
	public byte[] marshal(OutputStream out, byte[] buf) throws IOException {
		int n = 0;
		if (buf != null && buf.length != 0) try {
			n = marshal(buf, 0);
		} catch (BufferOverflowException e) {}
		if (n == 0) {
			buf = new byte[marshalFit()];
			n = marshal(buf, 0);
		}
		out.write(buf, 0, n);
		return buf;
	}

	/**
	 * Serializes the object.
	 * @param buf the data destination.
	 * @param offset the initial index for {@code buf}, inclusive.
	 * @return the final index for {@code buf}, exclusive.
	 * @throws BufferOverflowException when {@code buf} is too small.
	 * @throws IllegalStateException on an upper limit breach defined by {@link #colferSizeMax}.
	 */
	public int marshal(byte[] buf, int offset) {
		int i = offset;

		try {
			if (this.inner != null) {
				buf[i++] = (byte) 0;
				i = this.inner.marshal(buf, i);
			}

			buf[i++] = (byte) 0x7f;
			return i;
		} catch (ArrayIndexOutOfBoundsException e) {
			if (i - offset > EmbedO.colferSizeMax)
				throw new IllegalStateException(format("colfer: gen.EmbedO exceeds %d bytes", EmbedO.colferSizeMax));
			if (i > buf.length) throw new BufferOverflowException();
			throw e;
		}
	}

	/**
	 * Deserializes the object.
	 * @param buf the data source.
	 * @param offset the initial index for {@code buf}, inclusive.
	 * @return the final index for {@code buf}, exclusive.
	 * @throws BufferUnderflowException when {@code buf} is incomplete. (EOF)
	 * @throws SecurityException on an upper limit breach defined by {@link #colferSizeMax}.
	 * @throws InputMismatchException when the data does not match this object's schema.
	 */
	public int unmarshal(byte[] buf, int offset) {
		return unmarshal(buf, offset, buf.length);
	}

	/**
	 * Deserializes the object.
	 * @param buf the data source.
	 * @param offset the initial index for {@code buf}, inclusive.
	 * @param end the index limit for {@code buf}, exclusive.
	 * @return the final index for {@code buf}, exclusive.
	 * @throws BufferUnderflowException when {@code buf} is incomplete. (EOF)
	 * @throws SecurityException on an upper limit breach defined by {@link #colferSizeMax}.
	 * @throws InputMismatchException when the data does not match this object's schema.
	 */
	public int unmarshal(byte[] buf, int offset, int end) {
		if (end > buf.length) end = buf.length;
		int i = offset;

		try {
			byte header = buf[i++];

			if (header == (byte) 0) {
				this.inner = new O();
				i = this.inner.unmarshal(buf, i, end);
				header = buf[i++];
			}

			if (header != (byte) 0x7f)
				throw new InputMismatchException(format("colfer: unknown header at byte %d", i - 1));
		} finally {
			if (i > end && end - offset < EmbedO.colferSizeMax) throw new BufferUnderflowException();
			if (i < 0 || i - offset > EmbedO.colferSizeMax)
				throw new SecurityException(format("colfer: gen.EmbedO exceeds %d bytes", EmbedO.colferSizeMax));
			if (i > end) throw new BufferUnderflowException();
		}

		return i;
	}

	// {@link Serializable} version number.
	private static final long serialVersionUID = 1L;

	// {@link Serializable} Colfer extension.
	private void writeObject(ObjectOutputStream out) throws IOException {
		byte[] buf = new byte[marshalFit()];
		int n = marshal(buf, 0);
		out.writeInt(n);
		out.write(buf, 0, n);
	}

	// {@link Serializable} Colfer extension.
	private void readObject(ObjectInputStream in) throws ClassNotFoundException, IOException {
		init();

		int n = in.readInt();
		byte[] buf = new byte[n];
		in.readFully(buf);
		unmarshal(buf, 0);
	}

	// {@link Serializable} Colfer extension.
	private void readObjectNoData() throws ObjectStreamException {
		init();
	}

	/**
	 * Gets gen.EmbedO.inner.
	 * @return the value.
	 */
	public O getInner() {
		return this.inner;
	}

	/**
	 * Sets gen.EmbedO.inner.
	 * @param value the replacement.
	 */
	public void setInner(O value) {
		this.inner = value;
	}

	/**
	 * Sets gen.EmbedO.inner.
	 * @param value the replacement.
	 * @return {@code this}.
	 */
	public EmbedO withInner(O value) {
		this.inner = value;
		return this;
	}

	@Override
	public final int hashCode() {
		int h = 1;
		if (this.inner != null) h = 31 * h + this.inner.hashCode();
		return h;
	}

	@Override
	public final boolean equals(Object o) {
		return o instanceof EmbedO && equals((EmbedO) o);
	}

	public final boolean equals(EmbedO o) {
		if (o == null) return false;
		if (o == this) return true;

		return (this.inner == null ? o.inner == null : this.inner.equals(o.inner));
	}

}
